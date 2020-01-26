package application

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain/fme"
	"github.com/pkg/errors"
	"log"
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
		if err := f.lyricService.AddLyric(lb, fmeStruct.LyricDataPart.Colors); err != nil {
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
	fmeStruct.LyricDataPart.Colors = lyricColorPicker

	fmeBinary, err := fmeStruct.ExportBinary()
	if err != nil {
		return nil, errors.Wrap(err, "cannot export binary")
	}

	return fmeBinary, nil
}
