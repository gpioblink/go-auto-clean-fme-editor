package application

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
)


type FmeService struct {
	AddLyric error
}

// TODO: placeholder
func NewFmeService(fmeService FmeService) FmeService {
	return fmeService
}

//func (f FmeService) ImportFme(fme []byte) error {
//	err := CheckMagicValue(fme)
//	if err != nil {
//		return err
//	}
//
//	_, lyricOffset, _, err := getOffsets(fme)
//	if err != nil {
//		return err
//	}
//
//	lyricFme := fme[lyricOffset:]
//
//	for {
//		lyricSize := binary.LittleEndian.Uint32(fme[0:4])
//
//	}
//}

type FmeLyricPartHeader struct {
	flag uint16
	x uint16
	y uint16
	colorSelectBC byte
	colorSelectAC byte
	colorSelectBO byte
	colorSelectAO byte
}

type FmeLyricChar struct {
	fontCode byte
	char string
	width uint16
}

func CreateLyricPart(fme []byte, lyricPartOffset int) error {
	lastOffset := binary.LittleEndian.Uint16(fme[0:1])
	buf := bytes.NewReader(fme[lyricPartOffset:lyricPartOffset+int(lastOffset)])

	var header FmeLyricPartHeader
	if err := binary.Read(buf, binary.LittleEndian, &header); err !=nil {
	 	return err
	}

	var stringLength uint16
	if err := binary.Read(buf, binary.LittleEndian, &stringLength); err !=nil {
		return err
	}

	for i := 0; i < int(stringLength); i++ {

	}


}

type FmeColor struct {
	r int
	g int
	b int
}

func NewFmeColor(red int, green int, blue int) FmeColor {
	return FmeColor{red, green ,blue}
}

func CreateColorPalette(fme []byte, lyricOffset int) ([]FmeColor, error) {
	colorNum := 15
	var data uint16
	var fmeColors []FmeColor
	buf := bytes.NewReader(fme[lyricOffset:lyricOffset+colorNum*2])
	// lyricOffset+colorNum*2: end of the color header in the lyric data part
	for i := 0; i < colorNum; i++ {
		if err := binary.Read(buf, binary.LittleEndian, &data); err != nil {
			return nil, err
		}
		fmeColors = append(fmeColors, convertRGB555BinaryToRGB888Color(data))
	}
	return fmeColors, nil
}

func convertRGB555BinaryToRGB888Color(colorBin uint16) FmeColor {
	r := (colorBin & 0b0111110000000000) >> 10
	g := (colorBin & 0b0000001111100000) >> 5
	b := colorBin & 0b0000000000011111
	r = (r * 255) / 31
	g = (g * 255) / 31
	b = (b * 255) / 31
	return NewFmeColor(int(r), int(g), int(b))
}