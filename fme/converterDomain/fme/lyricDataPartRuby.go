package fme

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"math"
	"unicode/utf8"
)

type LyricRubyChar [2]byte

type LyricRuby struct {
	RubyCharCount           uint16
	RelativeHorizontalPoint uint16
	RubyChar                []LyricRubyChar
}

func (lr LyricRuby) GetRelativeHorizontalPoint() int {
	return int(lr.RelativeHorizontalPoint)
}

func (lr LyricRuby) GetRubyChar() string {
	var rubyBinary []byte
	for _, lrc := range lr.RubyChar {
		rubyBinary = append(rubyBinary, lrc[0])
		rubyBinary = append(rubyBinary, lrc[1])
	}
	utf8Str, _ := ConvertShiftJisToUTF8(rubyBinary)
	return utf8Str
}

func (lr LyricRuby) CalcByteSize() int {
	count := 0x02 + 0x02 // rubyCharCount, relativeHorizontalPoint
	count += len(lr.RubyChar) * 0x02
	return count
}

func NewLyricRuby(ruby string, horizontalPoint int) (*LyricRuby, error) {
	rc := utf8.RuneCountInString(ruby)
	rubyCountUint16 := uint16(rc)
	if !(0 < rc && rc < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	horizontalPointUint16 := uint16(horizontalPoint)
	if !(0 < horizontalPoint && horizontalPoint < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	var rubyBinary []LyricRubyChar
	for _, c := range ruby {
		rbc, err := ConvertUTF8CharToShiftJis(string(c))
		if err != nil {
			return nil, err
		}
		charByte := allocateTwoBytesSliceForTwoByte(rbc)
		rubyBinary = append(rubyBinary, charByte)
	}

	return &LyricRuby{rubyCountUint16, horizontalPointUint16, rubyBinary}, nil
}

func ConvertUTF8StringToShiftJis(s string) ([]byte, error) {
	t := japanese.ShiftJIS.NewEncoder()
	sjisStr, _, err := transform.String(t, s)
	if err != nil {
		return nil, err
	}

	b := []byte(sjisStr)

	return b, nil
}
