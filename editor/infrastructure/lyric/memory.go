package lyric

import (
	"errors"
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
)

var ErrIndexOutOfBounds = errors.New("index out of bounds")

type MemoryRepository struct {
	lyrics []lyric.Lyric
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]lyric.Lyric{}}
}

func (m *MemoryRepository) Save(lyricToSave *lyric.Lyric) error {
	m.lyrics = append(m.lyrics, *lyricToSave)
	return nil
}

func (m *MemoryRepository) Update(lyricToUpdate *lyric.Lyric, index int) error {
	if len(m.lyrics) < index {
		return ErrIndexOutOfBounds
	}
	m.lyrics[index] = *lyricToUpdate
	return nil
}

func (m *MemoryRepository) ListAllLyrics() ([]lyric.Lyric, error) {
	return m.lyrics, nil
}
