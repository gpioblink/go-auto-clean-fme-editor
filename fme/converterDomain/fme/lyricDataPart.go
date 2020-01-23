package fme

import (
	"bytes"
	"encoding/binary"
	"errors"
	errors2 "github.com/pkg/errors"
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
	RubyChar                []LyricRubyChar
}

type LyricRubyChar [2]byte

var ErrMultipleChar = errors.New("char must be a character")
var ErrBeyondBinary = errors.New("value beyond acceptable length")

func NewLyricBody(lyrics []LyricChar, ruby []LyricRuby) (*LyricBody, error) {
	// TODO: このキャストした数が上限値越えてないか見るのにもっといい方法ないかな？

	if !(0 < len(lyrics) && len(lyrics) < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}
	if !(0 < len(ruby) && len(ruby) < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	return &LyricBody{uint16(len(lyrics)), lyrics, uint16(len(ruby)), ruby}, nil
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

func allocateTwoBytesSliceForTwoByte(b []byte) [2]byte {
	var charByte [2]byte
	charByte[0] = b[1]
	charByte[1] = b[0]
	return charByte
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
		return nil, errors2.Wrap(err, "fail to read color")
	}

	var lyricBlocks []LyricBlock
	for {
		var lyricHeader LyricHeader
		err = binary.Read(buf, binary.LittleEndian, &lyricHeader)
		if err == io.EOF {
			break // break if I read all lyric
		} else if err != nil {
			return nil, errors2.Wrap(err, "fail to read header")
		}

		var lyricCount uint16
		err = binary.Read(buf, binary.LittleEndian, &lyricCount)
		if err != nil {
			return nil, errors2.Wrap(err, "fail to read lyricCount")
		}

		lyricString := make([]LyricChar, lyricCount)
		for i := 0; i < int(lyricCount); i++ {
			err = binary.Read(buf, binary.LittleEndian, &lyricString[i])
			if err != nil {
				return nil, errors2.Wrap(err, "fail to read lyricChar[%d]")
			}
		}

		var rubyCount uint16
		err = binary.Read(buf, binary.LittleEndian, &rubyCount)
		if err != nil {
			return nil, errors2.Wrap(err, "fail to read rubyCount")
		}

		var lyricRuby []LyricRuby
		for i := 0; i < int(rubyCount); i++ {
			var rubyCharCount uint16
			err = binary.Read(buf, binary.LittleEndian, &rubyCharCount)
			if err != nil {
				return nil, errors2.Wrap(err, "fail to read rubyCharCount")
			}

			var relativeHorizontalPoint uint16
			err = binary.Read(buf, binary.LittleEndian, &relativeHorizontalPoint)
			if err != nil {
				return nil, errors2.Wrap(err, "fail to read ruby horizontal point")
			}

			var rubyString []LyricRubyChar
			for j := 0; j < int(rubyCharCount); j++ {
				var rubyChar LyricRubyChar
				err = binary.Read(buf, binary.LittleEndian, &rubyChar)
				if err != nil {
					return nil, errors2.Wrap(err, "fail to read ruby char")
				}
				rubyString = append(rubyString, rubyChar)
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
		return nil, errors2.Wrap(err, "fail to read color")
	}

	for _, b := range d.LyricBlocks {
		// lyric header
		err = binary.Write(buf, binary.LittleEndian, b.LyricHeader)
		if err != nil {
			return nil, errors2.Wrap(err, "fail to write header")
		}

		// lyric body
		err = binary.Write(buf, binary.LittleEndian, b.LyricBody.LyricCount)
		if err != nil {
			return nil, errors2.Wrap(err, "fail to write lyricCount")
		}

		for _, c := range b.LyricBody.Lyrics {
			err = binary.Write(buf, binary.LittleEndian, c)
			if err != nil {
				return nil, errors2.Wrap(err, "fail to write lyricBody")
			}
		}

		err = binary.Write(buf, binary.LittleEndian, b.LyricBody.RubyCount)
		if err != nil {
			return nil, errors2.Wrap(err, "fail to write ruby count")
		}

		for _, r := range b.LyricBody.Ruby {
			err = binary.Write(buf, binary.LittleEndian, r.RubyCharCount)
			if err != nil {
				return nil, errors2.Wrap(err, "fail to write ruby char count")
			}
			err = binary.Write(buf, binary.LittleEndian, r.RelativeHorizontalPoint)
			if err != nil {
				return nil, errors2.Wrap(err, "fail to write ruby horizontal point")
			}
			for _, rb := range r.RubyChar {
				err = binary.Write(buf, binary.LittleEndian, rb)
				if err != nil {
					return nil, errors2.Wrap(err, "fail to write ruby char")
				}
			}
		}
	}
	return buf.Bytes(), nil
}
