package fme

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"math"
	"unicode/utf8"
)

type LyricDataPart struct {
	LyricColorPicker
	LyricBlocks []LyricBlock
}

type LyricBlock struct {
	LyricHeader
	LyricBody
}

type LyricColorPicker struct {
	Colors [15]uint16
}

type LyricHeader struct {
	LyricDataSize uint16
	Flag          uint16
	X             uint16
	Y             uint16
	ColorSelectBC byte
	ColorSelectAC byte
	ColorSelectBO byte
	ColorSelectAO byte
}

type LyricBody struct {
	LyricCount uint16
	Lyrics     []LyricChar
	RubyCount  uint16
	Ruby       []LyricRuby
}

type LyricChar struct {
	FontCode byte
	Char     [2]byte
	Width    uint16
}

type LyricRuby struct {
	RubyCharCount           uint16
	RelativeHorizontalPoint uint16
	RubyChar                []byte
}

var ErrMultipleChar = errors.New("char must be a character")
var ErrBeyondBinary = errors.New("width beyond acceptable length")

func NewLyricChar(char string, width int) (*LyricChar, error) {
	fontCode := byte(0x00) // shift_jis

	var charByte [2]byte
	b, err := ConvertUTF8CharToShiftJis(char)
	if err != nil {
		return nil, err
	}

	charByte[0] = b[1]
	charByte[1] = b[0]

	widthTime := uint16(width)
	if !(0 < width && width < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	return &LyricChar{fontCode, charByte, widthTime}, nil
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

func NewLyricDataPartFromBinary(fme []byte) (*LyricDataPart, error) {
	// TODO: 各構造体をNew*FromBinaryに分けて残す

	buf := bytes.NewBuffer(fme)

	var lyricColorPicker LyricColorPicker
	err := binary.Read(buf, binary.LittleEndian, &lyricColorPicker)
	if err != nil {
		return nil, err
	}

	var lyricBlocks []LyricBlock
	for {
		var lyricHeader LyricHeader
		err = binary.Read(buf, binary.LittleEndian, &lyricHeader)
		if err == io.EOF {
			break // break if I read all lyric
		} else if err != nil {
			return nil, err
		}

		var lyricCount uint16
		err = binary.Read(buf, binary.LittleEndian, &lyricCount)
		if err != nil {
			return nil, err
		}

		lyricString := make([]LyricChar, lyricCount)
		for i := 0; i < int(lyricCount); i++ {
			err = binary.Read(buf, binary.LittleEndian, &lyricString[i])
			if err != nil {
				return nil, err
			}
		}

		var rubyCount uint16
		err = binary.Read(buf, binary.LittleEndian, &rubyCount)
		if err != nil {
			return nil, err
		}

		var lyricRuby []LyricRuby
		for i := 0; i < int(rubyCount); i++ {
			var rubyCharCount uint16
			err = binary.Read(buf, binary.LittleEndian, &rubyCharCount)
			if err != nil {
				return nil, err
			}

			var relativeHorizontalPoint uint16
			err = binary.Read(buf, binary.LittleEndian, &relativeHorizontalPoint)
			if err != nil {
				return nil, err
			}

			rubyString := make([]byte, rubyCharCount*2) // shift-jis needs 2bytes par a character
			err = binary.Read(buf, binary.LittleEndian, &rubyString)
			if err != nil {
				return nil, err
			}

			lyricRuby = append(lyricRuby, LyricRuby{
				rubyCharCount, relativeHorizontalPoint, rubyString,
			})
		}

		lyricBody := LyricBody{lyricCount, lyricString, rubyCount, lyricRuby}

		lyricBlocks = append(lyricBlocks, LyricBlock{lyricHeader, lyricBody})
	}

	return &LyricDataPart{lyricColorPicker, lyricBlocks}, nil
}

func (d *LyricDataPart) ExportBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, d.Colors)
	if err != nil {
		return nil, err
	}

	for _, b := range d.LyricBlocks {
		// lyric header
		err = binary.Write(buf, binary.LittleEndian, b.LyricHeader)
		if err != nil {
			return nil, err
		}

		// lyric body
		err = binary.Write(buf, binary.LittleEndian, b.LyricBody.LyricCount)
		if err != nil {
			return nil, err
		}

		for _, c := range b.LyricBody.Lyrics {
			err = binary.Write(buf, binary.LittleEndian, c)
			if err != nil {
				return nil, err
			}
		}

		err = binary.Write(buf, binary.LittleEndian, b.LyricBody.RubyCount)
		if err != nil {
			return nil, err
		}

		for _, r := range b.LyricBody.Ruby {
			err = binary.Write(buf, binary.LittleEndian, r.RubyCharCount)
			if err != nil {
				return nil, err
			}
			err = binary.Write(buf, binary.LittleEndian, r.RelativeHorizontalPoint)
			if err != nil {
				return nil, err
			}
			for _, rb := range r.RubyChar {
				err = binary.Write(buf, binary.LittleEndian, rb)
			}
		}
	}
	return buf.Bytes(), nil
}
