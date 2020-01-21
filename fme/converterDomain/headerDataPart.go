package fme

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type HeaderDataPart struct {
	FileId                    [6]byte
	InformationDataPartOffset uint32
	LyricOffset               uint32
	TimingOffset              uint32
}

func NewHeaderDataPartFromBinary(fme []byte) (*HeaderDataPart, error) {
	buf := bytes.NewBuffer(fme)

	err := CheckMagicValue(fme)
	if err != nil {
		return nil, err
	}

	var header HeaderDataPart
	err = binary.Read(buf, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	return &header, nil
}

func (d *HeaderDataPart) ExportBinary() ([]byte, error) {
	bufHeader := new(bytes.Buffer)
	err := binary.Write(bufHeader, binary.LittleEndian, d)
	if err != nil {
		return nil, err
	}
	return bufHeader.Bytes(), nil
}

func (d *HeaderDataPart) GetOffsets() (infoDataOffset uint32, lyricOffset uint32, timingOffset uint32) {
	return d.InformationDataPartOffset, d.LyricOffset, d.TimingOffset
}

var ErrInvalidMagicNumber = errors.New("invalid magic")

func CheckMagicValue(fme []byte) error {
	magic := []byte{0x4A, 0x4F, 0x59, 0x2D, 0x30, 0x32}
	if bytes.Equal(magic, fme[:6]) == false {
		return ErrInvalidMagicNumber
	}
	return nil
}

func GetOffsets(fme []byte) (infoDataOffset uint32, lyricOffset uint32, timingOffset uint32, err error) {
	infoDataOffset = binary.LittleEndian.Uint32(fme[6:10])
	lyricOffset = binary.LittleEndian.Uint32(fme[10:14])
	timingOffset = binary.LittleEndian.Uint32(fme[14:18])
	return infoDataOffset, lyricOffset, timingOffset, nil
}
