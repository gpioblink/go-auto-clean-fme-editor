package lyric

import (
	"errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"unicode/utf8"
)

var ErrMultipleCharactersInChar = errors.New("char must have 1 character")
var ErrConvertShiftJIS = errors.New("cannot use a char that is not available in shift_jis")
var ErrInvalidLength = errors.New("length is positive value")

type LyricChar struct {
	char   string
	length int
}

func (lc LyricChar) Char() string {
	return lc.char
}

func (lc LyricChar) Length() int {
	return lc.length
}

func NewLyricChar(char string, length int) (*LyricChar, error) {
	if utf8.RuneCountInString(char) != 1 {
		return nil, ErrMultipleCharactersInChar
	}

	if length < 0 {
		return nil, ErrInvalidLength
	}

	if _, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), char); err != nil {
		return nil, ErrConvertShiftJIS
	}

	return &LyricChar{char, length}, nil
}
