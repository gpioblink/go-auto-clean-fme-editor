package fme

import (
	"bytes"
	"encoding/binary"
)

type HeaderDataPart struct {
	FileId                    [6]byte
	InformationDataPartOffset uint32
	LyricOffset               uint32
	TimingOffset              uint32
}

func NewHeaderDataPartFromBinary(fme []byte) (*HeaderDataPart, error) {
	buf := bytes.NewBuffer(fme)

	var header HeaderDataPart
	err := binary.Read(buf, binary.LittleEndian, &header)
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
