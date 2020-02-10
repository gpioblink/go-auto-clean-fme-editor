package fme

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	errors2 "github.com/pkg/errors"
)

const LyricHeaderSize = 0xc

var (
	ErrMultipleChar  = errors.New("char must be a character")
	ErrBeyondBinary  = errors.New("value beyond acceptable length")
	ErrColorNotFound = errors.New("color not found")
)

type LyricDataPart struct {
	Colors      LyricColorPicker
	LyricBlocks []LyricBlock
}

func NewLyricDataPart(colorPicker LyricColorPicker, lyricBlocks []LyricBlock) (*LyricDataPart, error) {
	return &LyricDataPart{colorPicker, lyricBlocks}, nil
}

// 仕様がわかんないからあれだけど雰囲気で書いていく
// とりあえず、それぞれのバイナリのを読むところを名前をつけて関数に切った方が後で読んだときにここは何を読んでるかわかりやすいと思いました！
// 所々、コメントで他の関数はその辺担保しているところも見られたので、コメントでもいいかも
func NewLyricDataPartFromBinary(fme []byte) (*LyricDataPart, error) {
	// TODO: 各構造体をNew*FromBinaryに分けて残す

	// 僕だったらbytes.NewReader
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
