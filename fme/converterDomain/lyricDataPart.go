package fme

import (
	"bytes"
	"encoding/binary"
	"io"
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

func NewLyricDataPartFromBinary(fme []byte) (*LyricDataPart, error) {
	// TODO: complex code. need some cleaning

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
