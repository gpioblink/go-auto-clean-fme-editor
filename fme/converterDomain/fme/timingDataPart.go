package fme

import (
	"bytes"
	"encoding/binary"
	"io"
)

type TimingDataPart struct {
	timingData []TimingData
}

type TimingData struct {
	EventTime uint32
	DataSize  byte
	EventType byte
	EventData []byte
}

func NewTimingDataPartFromBinary(fme []byte) (*TimingDataPart, error) {
	buf := bytes.NewBuffer(fme)

	var timingDataArray []TimingData

	for {
		var eventTime uint32
		err := binary.Read(buf, binary.LittleEndian, &eventTime)
		if err == io.EOF {
			break // if the file was end, break
		} else if err != nil {
			return nil, err
		}

		var dataSize byte
		err = binary.Read(buf, binary.LittleEndian, &dataSize)
		if err != nil {
			return nil, err
		}

		var eventType byte
		err = binary.Read(buf, binary.LittleEndian, &eventType)
		if err != nil {
			return nil, err
		}

		eventData := make([]byte, dataSize-1) // dataSize-1: dataSize includes eventType(1byte)
		err = binary.Read(buf, binary.LittleEndian, &eventData)
		if err != nil {
			return nil, err
		}

		timingData := TimingData{
			eventTime, dataSize, eventType, eventData,
		}
		timingDataArray = append(timingDataArray, timingData)
	}

	return &TimingDataPart{timingDataArray}, nil
}

func (t TimingDataPart) ExportBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	for _, td := range t.timingData {
		err := binary.Write(buf, binary.LittleEndian, td.EventTime)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buf, binary.LittleEndian, td.DataSize)
		if err != nil {
			return nil, err
		}
		err = binary.Write(buf, binary.LittleEndian, td.EventType)
		if err != nil {
			return nil, err
		}

		for _, tdd := range td.EventData {
			err = binary.Write(buf, binary.LittleEndian, tdd)
			if err != nil {
				return nil, err
			}
		}
	}
	return buf.Bytes(), nil
}
