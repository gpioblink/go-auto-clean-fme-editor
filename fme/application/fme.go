package application

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain/fme"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type lyricService interface {
	AddLyric(block fme.LyricBlock, colorPicker fme.LyricColorPicker) error
	ListLyrics() (blocks []fme.LyricBlock, colorPicker fme.LyricColorPicker, err error)
	// TODO: need ClearLyrics to use continuously
}

type FmeService struct {
	lyricService lyricService

	fmeRepository fme.Repository
}

func NewFmeService(lyricService lyricService, lyricRepository fme.Repository) FmeService {
	return FmeService{lyricService, lyricRepository}
}

func (f FmeService) ImportFme(fmeByte []byte) error {
	fmeStruct, err := fme.NewFmeFromBinary(fmeByte)
	if err != nil {
		return errors.Wrap(err, "invalid or unknown binary")
	}

	if err := f.fmeRepository.Save(fmeStruct); err != nil {
		return errors.Wrap(err, "cannot save fme to storage")
	}

	for _, lb := range fmeStruct.LyricDataPart.LyricBlocks {
		if err := f.lyricService.AddLyric(lb, fmeStruct.LyricDataPart.LyricColorPicker); err != nil {
			return errors.Wrap(err, "cannot add lyric")
		}
	}

	log.Printf("fme file #{fme.fmeStruct} imported")

	return nil
}

func (f FmeService) ExportFme() ([]byte, error) {
	fmeStruct, err := f.fmeRepository.Get()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get fme struct")
	}

	lyricBlocks, lyricColorPicker, err := f.lyricService.ListLyrics()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get edited lyric")
	}

	// TODO: make appropriate functions
	fmeStruct.LyricDataPart.LyricBlocks = lyricBlocks
	fmeStruct.LyricDataPart.LyricColorPicker = lyricColorPicker

	fmeBinary, err := f.ExportFme()
	if err != nil {
		return nil, errors.Wrap(err, "cannot export binary")
	}

	return fmeBinary, nil
}

//func convertRGB555BinaryToRGB888Color(colorBin uint16) FmeColor {
//	r := (colorBin & 0b0111110000000000) >> 10
//	g := (colorBin & 0b0000001111100000) >> 5
//	b := colorBin & 0b0000000000011111
//	r = (r * 255) / 31
//	g = (g * 255) / 31
//	b = (b * 255) / 31
//	return NewFmeColor(int(r), int(g), int(b))
//}

//func TestCreateColorPalette(t *testing.T) {
//	fme := decodeTestBytes()
//	fmeColors, err := application.CreateColorPalette(fme, 0x77)
//	assert.NoError(t, err)
//
//	fmeExpectedColors := []application.FmeColor{
//		application.NewFmeColor(0x08, 0x08, 0x08),
//		application.NewFmeColor(0xff, 0xff, 0xff),
//		application.NewFmeColor(0xff, 0xff, 0x39),
//		application.NewFmeColor(0xff, 0x29, 0xff),
//		application.NewFmeColor(0xff, 0x94, 0x00),
//		application.NewFmeColor(0xff, 0x29, 0xff),
//		application.NewFmeColor(0x00, 0xf6, 0x00),
//		application.NewFmeColor(0x00, 0xf6, 0xff),
//		application.NewFmeColor(0x00, 0x39, 0x7b),
//		application.NewFmeColor(0x00, 0x52, 0x00),
//		application.NewFmeColor(0xb4, 0x00, 0x00),
//		application.NewFmeColor(0x8b, 0x00, 0x8b),
//		application.NewFmeColor(0x6a, 0x08, 0x00),
//		application.NewFmeColor(0x00, 0x00, 0x00),
//		application.NewFmeColor(0x00, 0x00, 0x00),
//	}
//
//	assert.EqualValues(t, fmeExpectedColors, fmeColors)
//}
