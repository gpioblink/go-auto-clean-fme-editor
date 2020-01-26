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
	Colors      LyricColorPicker
	LyricBlocks []LyricBlock
}

type LyricBlock struct {
	LyricHeader
	LyricBody
}

func (lb LyricBlock) CalcByteSize() int {
	return 0xc + lb.LyricBody.CalcByteSize()
}

type LyricColorPicker struct { // fmeで入力される色はいつも一緒なので変数名にしちゃいました
	DarkGray         uint16
	White            uint16
	Yellow           uint16
	Pink             uint16
	Orange           uint16
	Pink2            uint16
	LightGreen       uint16
	LightBlue        uint16
	DarkBlue         uint16
	DarkGreen        uint16
	GeneralMotorsRed uint16
	Purple           uint16
	Brown            uint16
	Custom1          uint16
	Custom2          uint16
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

func (lb LyricBody) CalcByteSize() int {
	count := 0x2 + 0x2 // LyricCount, RubyCount
	for _, l := range lb.Lyrics {
		count += l.CalcByteSize()
	}
	for _, r := range lb.Ruby {
		count += r.CalcByteSize()
	}
	return count
}

type LyricChar struct {
	FontCode byte
	Char     [2]byte
	Width    uint16
}

func (lc LyricChar) CalcByteSize() int {
	return 0x05
}

type LyricRuby struct {
	RubyCharCount           uint16
	RelativeHorizontalPoint uint16
	RubyChar                []LyricRubyChar
}

func (lr LyricRuby) CalcByteSize() int {
	count := 0x02 + 0x02 // rubyCharCount, relativeHorizontalPoint
	count += len(lr.RubyChar) * 0x02
	return count
}

type LyricRubyChar [2]byte

var ErrMultipleChar = errors.New("char must be a character")
var ErrBeyondBinary = errors.New("value beyond acceptable length")
var ErrColorNotFound = errors.New("color not found")

var StandardColorPicker = LyricColorPicker{0x0421, 0x7fff, 0x7fe7, 0x7cbf, 0x7e40,
	0x7cbf, 0x03c0, 0x03df, 0x00ef, 0x0140, 0x5800,
	0x4411, 0x3420, 0x0000, 0x0000,
}

func NewLyricDataPart(colorPicker LyricColorPicker, lyricBlocks []LyricBlock) (*LyricDataPart, error) {
	return &LyricDataPart{colorPicker, lyricBlocks}, nil
}

func NewLyricBlock(header LyricHeader, body LyricBody) (*LyricBlock, error) {
	return &LyricBlock{header, body}, nil
}

func (cp *LyricColorPicker) FindColorIndex(rgb555 uint16) (byte, error) {
	picker := cp
	colorBin := []uint16{picker.DarkGray, picker.White, picker.Yellow, picker.Pink, picker.Orange,
		picker.Pink2, picker.LightGreen, picker.LightBlue, picker.DarkBlue, picker.DarkGreen, picker.GeneralMotorsRed,
		picker.Purple, picker.Brown, picker.Custom1, picker.Custom2}

	for i, v := range colorBin {
		if rgb555 == v {
			return byte(i), nil
		}
	}
	return 255, ErrColorNotFound
}

func NewLyricHeaderWithStandardColorPicker(lyricBodySize int, x int, y int, beforeCharColor Color, afterCharColor Color,
	beforeOutlineColor Color, afterOutlineColor Color) (*LyricHeader, error) {

	picker := StandardColorPicker

	bc, err := picker.FindColorIndex(beforeCharColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	ac, err := picker.FindColorIndex(afterCharColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	bo, err := picker.FindColorIndex(beforeOutlineColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	ao, err := picker.FindColorIndex(afterOutlineColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}

	return &LyricHeader{uint16(lyricBodySize + 0x09), 0x00, uint16(x), uint16(y), bc, ac, bo, ao}, nil
}

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
