package informationDataPart

import (
	"bytes"
	"encoding/binary"
)

type InformationDataPart struct {
	Header
	Body
}

type Header struct {
	MusicPartsType       uint16
	MusicNameOffset      uint16
	SingerNameOffset     uint16
	LyricWriterOffset    uint16
	MusicWriterOffset    uint16
	MusicNameKanaOffset  uint16
	SingerNameKanaOffset uint16
	JasracCodeOffset     uint16
	LyricHeadOffset      uint16
	FileId               uint16
	VocalTracks          uint32
	RhythmTracks         uint32
}

type Body struct {
	MusicName      []byte
	SingerName     []byte
	LyricWriter    []byte
	MusicWriter    []byte
	MusicNameKana  []byte
	SingerNameKana []byte
	JasracCode     []byte
	LyricHead      []byte
}

func NewInformationDataPartFromBinary(fme []byte) (*InformationDataPart, error) {
	buf := bytes.NewBuffer(fme)

	var header Header
	err := binary.Read(buf, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	bodyArray := bytes.SplitAfter(fme[header.MusicNameOffset:], []byte{0x00})
	body := Body{
		bodyArray[0],
		bodyArray[1],
		bodyArray[2],
		bodyArray[3],
		bodyArray[4],
		bodyArray[5],
		bodyArray[6],
		bodyArray[7],
	}

	return &InformationDataPart{header, body}, nil
}

func (d *InformationDataPart) ExportBinary() ([]byte, error) {
	bufHeader := new(bytes.Buffer)

	err := binary.Write(bufHeader, binary.LittleEndian, d.Header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.MusicName)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.SingerName)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.LyricWriter)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.MusicWriter)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.MusicNameKana)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.SingerNameKana)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.JasracCode)
	if err != nil {
		return nil, err
	}
	err = binary.Write(bufHeader, binary.LittleEndian, d.Body.LyricHead)
	if err != nil {
		return nil, err
	}
	// bytesData := append(bufHeader.Bytes(), bufBody.Bytes()...)
	return bufHeader.Bytes(), nil
}
