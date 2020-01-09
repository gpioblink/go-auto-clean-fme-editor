package lyric

import (
	"errors"
)

var ErrEmptyLyric = errors.New("empty lyrics")

type LyricString []LyricChar

type Lyric struct {
	point  Point
	colors ColorPicker
	lyric  LyricString
}

func (l Lyric) Point() Point {
	return l.point
}

func (l Lyric) Colors() ColorPicker {
	return l.colors
}

func (l Lyric) Lyric() LyricString {
	return l.lyric
}

func NewLyric(point Point, colors ColorPicker, lyric LyricString) (*Lyric, error) {
	if len(lyric) == 0 {
		return nil, ErrEmptyLyric
	}
	return &Lyric{point, colors, lyric}, nil
}
