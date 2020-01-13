package application

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/pkg/errors"
)

type lyricReadModel interface {
	ListAllLyrics() ([]lyric.Lyric, error)
}

type LyricService struct {
	lyricReadModel  lyricReadModel
	lyricRepository lyric.Repository
}

func NewLyricService(lyricReadModel lyricReadModel, lyricRepository lyric.Repository) LyricService {
	return LyricService{lyricReadModel, lyricRepository}
}

func (s LyricService) ListLyrics() ([]lyric.Lyric, error) {
	return s.lyricReadModel.ListAllLyrics()
}

func (s LyricService) AddLyric(lyric lyric.Lyric) error {
	err := s.lyricRepository.Save(&lyric)
	if err != nil {
		return errors.Wrap(err, "cannot save lyric")
	}
	return nil
}

func (s LyricService) EditLyric(cmd EditLyricCommand) error {

	// get original lyric
	originalLyric, err := s.lyricRepository.ByIndex(cmd.Index)
	if err != nil {
		return errors.Wrap(err, "cannot get original lyric")
	}

	// create lyricString from Lyric
	newLyricString := lyric.LyricString{}
	for _, l := range cmd.Lyrics {
		lc, err := lyric.NewLyricChar(l.LyricChar, l.Length, l.Furigana)
		if err != nil {
			return errors.Wrap(err, "creating lyricString failed")
		}

		newLyricString = append(newLyricString, *lc)
	}

	// insert lyricString into the original and create a new lyric
	newLyric, err := lyric.NewLyric(originalLyric.Point(), originalLyric.Colors(), newLyricString)
	if err != nil {
		return errors.Wrap(err, "merging lyricString failed")
	}

	err = s.lyricRepository.Update(newLyric, cmd.Index)
	if err != nil {
		return errors.Wrap(err, "cannot update lyric")
	}

	return nil
}

type EditLyricCommand struct {
	Lyrics []EditLyricLyric
	Index  int
}

type EditLyricLyric struct {
	Furigana  string
	Length    int
	LyricChar string
}
