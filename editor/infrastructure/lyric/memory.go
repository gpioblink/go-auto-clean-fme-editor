package lyric

import (
	"errors"
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
)

var ErrIndexNotFound = errors.New("index not found")

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
	if len(m.lyrics) <= index {
		return ErrIndexNotFound
	}
	m.lyrics[index] = *lyricToUpdate
	return nil
}

func (m MemoryRepository) ByIndex(index int) (*lyric.Lyric, error) {
	if len(m.lyrics) <= index {
		return nil, ErrIndexNotFound
	}

	return &m.lyrics[index], nil
}

func (m *MemoryRepository) ListAllLyrics() ([]lyric.Lyric, error) {
	return m.lyrics, nil
}
