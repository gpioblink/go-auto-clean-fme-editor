package fme

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"math"
	"unicode/utf8"
)

type LyricChar struct {
	FontCode byte
	Char     [2]byte
	Width    uint16
}

func (lc LyricChar) GetChar() string {
	str, _ := ConvertShiftJisToUTF8([]byte{lc.Char[0], lc.Char[1]})
	return str
}

func (lc LyricChar) GetWidth() int {
	return int(lc.Width)
}

func (lc LyricChar) CalcByteSize() int {
	return 0x05
}

func NewLyricChar(char string, width int) (*LyricChar, error) {
	fontCode := byte(0x00) // shift_jis

	b, err := ConvertUTF8CharToShiftJis(char)
	if err != nil {
		return nil, err
	}

	charByte := allocateTwoBytesSliceForTwoByte(b)

	widthTime := uint16(width)
	if !(0 < width && width < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	return &LyricChar{fontCode, charByte, widthTime}, nil
}

func allocateTwoBytesSliceForTwoByte(b []byte) [2]byte {
	var charByte [2]byte
	charByte[0] = b[1]
	charByte[1] = b[0]
	return charByte
}

func ConvertShiftJisToUTF8(byte []byte) (string, error) {
	t := japanese.ShiftJIS.NewDecoder()
	utf8Str, _, err := transform.Bytes(t, byte)
	if err != nil {
		return "", err
	}

	s := string(utf8Str)

	return s, nil
}

func ConvertUTF8CharToShiftJis(s string) ([]byte, error) {
	// TODO: 英字(半角)の扱いを要調査。変換後が1バイトだった場合の扱いが不明
	if utf8.RuneCountInString(s) != 1 {
		return nil, ErrMultipleChar
	}

	b, err := ConvertUTF8StringToShiftJis(s)
	if err != nil {
		return nil, err
	}

	return b, nil
}
