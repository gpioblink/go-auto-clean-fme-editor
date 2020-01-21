package fme

import (
	"bytes"
	"encoding/binary"
)

type InformationDataPart struct {
	InformationDataPartHeader
	InformationDataPartBody
}

type InformationDataPartHeader struct {
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

type InformationDataPartBody struct {
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

	var header InformationDataPartHeader
	err := binary.Read(buf, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	bodyArray := bytes.SplitAfter(fme[header.MusicNameOffset:], []byte{0x00})
	body := InformationDataPartBody{
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
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, d.InformationDataPartHeader)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.MusicName)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.SingerName)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.LyricWriter)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.MusicWriter)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.MusicNameKana)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.SingerNameKana)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.JasracCode)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, d.InformationDataPartBody.LyricHead)
	if err != nil {
		return nil, err
	}
	// bytesData := append(bufHeader.Bytes(), bufBody.Bytes()...)
	return buf.Bytes(), nil
}
