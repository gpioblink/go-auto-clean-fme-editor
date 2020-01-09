package application

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/pkg/errors"
)

type lyricReadModel interface {
	ListAllProducts() ([]lyric.Lyric, error)
}

type LyricService struct {
	lyricReadModel  lyricReadModel
	lyricRepository lyric.Repository
}

func NewLyricService(lyricReadModel lyricReadModel, lyricRepository lyric.Repository) LyricService {
	return LyricService{lyricReadModel, lyricRepository}
}

func (s LyricService) ListLyrics() ([]lyric.Lyric, error) {
	return s.lyricReadModel.ListAllProducts()
}

func (s LyricService) AddLyric(lyric lyric.Lyric) error {
	err := s.lyricRepository.Save(&lyric)
	if err != nil {
		return errors.Wrap(err, "cannot save lyric")
	}
	return nil
}

func (s LyricService) EditLyric(lyric lyric.Lyric, index int) error {
	err := s.lyricRepository.Update(&lyric, index)
	if err != nil {
		return errors.Wrap(err, "cannot update lyric")
	}
	return nil
}
