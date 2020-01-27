package fme_test

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/converterDomain/fme"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLyricHeaderWithStandardColorPicker(t *testing.T) {
	testCases := []struct {
		TestName       string
		ExpectedErr    bool
		LyricBodySize  int
		X              int
		Y              int
		bcColor        fme.Color
		acColor        fme.Color
		boColor        fme.Color
		aoColor        fme.Color
		ExpectedBinary []byte
	}{
		{
			TestName:       "基本",
			ExpectedErr:    false,
			LyricBodySize:  0x45,
			X:              0x114,
			Y:              0x17f,
			bcColor:        *fme.NewColor(0x7fff),
			acColor:        *fme.NewColor(0x5800),
			boColor:        *fme.NewColor(0x0421),
			aoColor:        *fme.NewColor(0x7fff),
			ExpectedBinary: []byte{0x51, 0x00, 0x00, 0x00, 0x14, 0x01, 0x7f, 0x01, 0x01, 0x0a, 0x00, 0x01},
		},
		{
			TestName:       "全部ff",
			ExpectedErr:    true,
			LyricBodySize:  0x48,
			X:              0x114,
			Y:              0x17f,
			bcColor:        *fme.NewColor(0xffff),
			acColor:        *fme.NewColor(0xffff),
			boColor:        *fme.NewColor(0xffff),
			aoColor:        *fme.NewColor(0xffff),
			ExpectedBinary: []byte{0x51, 0x00, 0x00, 0x00, 0x14, 0x01, 0x7f, 0x01, 0x01, 0x0a, 0x00, 0x01},
		},
		{
			TestName:       "存在しない色",
			ExpectedErr:    true,
			LyricBodySize:  0x48,
			X:              0x114,
			Y:              0x17f,
			bcColor:        *fme.NewColor(0x7b6f),
			acColor:        *fme.NewColor(0x4532),
			boColor:        *fme.NewColor(0x6543),
			aoColor:        *fme.NewColor(0x65ab),
			ExpectedBinary: []byte{0x51, 0x00, 0x00, 0x00, 0x14, 0x01, 0x7f, 0x01, 0x01, 0x0a, 0x00, 0x01},
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			lyricHeader, err := fme.NewLyricHeaderWithStandardColorPicker(
				c.LyricBodySize, c.X, c.Y, c.bcColor, c.acColor, c.boColor, c.aoColor)

			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.LittleEndian, lyricHeader)
				assert.NoError(t, err)
				assert.EqualValues(t, c.ExpectedBinary, buf.Bytes())
			}

		})
	}
}

func TestNewLyricChar(t *testing.T) {
	testCases := []struct {
		TestName       string
		ExpectedErr    bool
		char           string
		width          int
		ExpectedBinary []byte
	}{
		{
			TestName:       "基本",
			ExpectedErr:    false,
			char:           "君",
			width:          48,
			ExpectedBinary: []byte{0x00, 0x4e, 0x8c, 0x30, 0x00},
		},
		{
			TestName:       "ひらがな",
			ExpectedErr:    false,
			char:           "が",
			width:          48,
			ExpectedBinary: []byte{0x00, 0xaa, 0x82, 0x30, 0x00},
		},
		{
			TestName:       "半角英字",
			ExpectedErr:    false,
			char:           "o",
			width:          26,
			ExpectedBinary: []byte{0x00, 0x6f, 0x00, 0x1a, 0x00},
		},
		{
			TestName:       "複数文字",
			ExpectedErr:    true,
			char:           "あいうえお",
			width:          23,
			ExpectedBinary: nil,
		},
		{
			TestName:    "ShiftJIS非対応文字",
			ExpectedErr: true,
			char:        "아",
			width:       48,
		},
		{
			TestName:    "大きすぎるwidth",
			ExpectedErr: true,
			char:        "あ",
			width:       999999999999999999,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			lyricChar, err := fme.NewLyricChar(c.char, c.width)

			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.LittleEndian, lyricChar)
				assert.NoError(t, err)
				assert.EqualValues(t, c.ExpectedBinary, buf.Bytes())

				// test getter
				assert.EqualValues(t, c.char, lyricChar.GetChar())
				assert.EqualValues(t, c.width, lyricChar.GetWidth())
			}

		})
	}
}

func TestNewLyricRuby(t *testing.T) {
	testCases := []struct {
		TestName       string
		ExpectedErr    bool
		ruby           string
		rubyPoint      int
		ExpectedBinary []byte
	}{
		{
			TestName:       "2文字",
			ExpectedErr:    false,
			ruby:           "きみ",
			rubyPoint:      0x0,
			ExpectedBinary: []byte{0x02, 0x00, 0x00, 0x00, 0xab, 0x82, 0xdd, 0x82},
		},
		{
			TestName:       "1文字",
			ExpectedErr:    false,
			ruby:           "ち",
			rubyPoint:      0xc,
			ExpectedBinary: []byte{0x01, 0x00, 0x0c, 0x00, 0xbf, 0x82},
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			lyricRuby, err := fme.NewLyricRuby(c.ruby, c.rubyPoint)

			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.LittleEndian, lyricRuby.RubyCharCount)
				assert.NoError(t, err)
				err = binary.Write(buf, binary.LittleEndian, lyricRuby.RelativeHorizontalPoint)
				assert.NoError(t, err)
				err = binary.Write(buf, binary.LittleEndian, lyricRuby.RubyChar)
				assert.NoError(t, err)
				assert.EqualValues(t, c.ExpectedBinary, buf.Bytes())

				// test getter
				assert.EqualValues(t, c.ruby, lyricRuby.GetRubyChar())
				assert.EqualValues(t, c.rubyPoint, lyricRuby.GetRelativeHorizontalPoint())
			}

		})
	}
}

func TestConvertUTF8CharToShiftJisAndReverse(t *testing.T) {
	testCases := []struct {
		TestName    string
		ExpectedErr bool
		utfString   string
		sjisByte    []byte
	}{
		{
			TestName:    "基本",
			ExpectedErr: false,
			utfString:   "あ",
			sjisByte:    []byte{0x82, 0xa0},
		},
		{
			TestName:    "漢字",
			ExpectedErr: false,
			utfString:   "一",
			sjisByte:    []byte{0x88, 0xea},
		},
		{
			TestName:    "複数文字列",
			ExpectedErr: true,
			utfString:   "一二三",
			sjisByte:    nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			sjisByte, err := fme.ConvertUTF8CharToShiftJis(c.utfString)
			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, c.sjisByte, sjisByte)

				utf8Str, err := fme.ConvertShiftJisToUTF8(sjisByte)
				assert.NoError(t, err)
				assert.EqualValues(t, c.utfString, utf8Str)
			}
		})
	}
}

func TestLyricDataPart_ExportBinary(t *testing.T) {
	fmeData := decodeLyricDataTestBytes()
	lyricData, err := fme.NewLyricDataPartFromBinary(fmeData)
	assert.NoError(t, err)

	fmeOut, err := lyricData.ExportBinary()
	assert.NoError(t, err)

	assert.EqualValues(t, fmeData, fmeOut)
}

func TestLyricBody_CalcByteSize(t *testing.T) {
	fmeData := decodeLyricDataTestBytes()
	lyricData, err := fme.NewLyricDataPartFromBinary(fmeData)
	assert.NoError(t, err)

	for _, lb := range lyricData.LyricBlocks {
		originalDataSize := lb.LyricHeader.LyricDataSize
		calculatedDataSize := lb.CalcByteSize()
		assert.EqualValues(t, originalDataSize, calculatedDataSize)
	}
}

func decodeLyricDataTestBytes() []byte {
	kimigayoBase64 := "IQT/f+d/v3xAfr98wAPfA+8AQAEAWBFEIDQAAAAAMgAAAHQAIQEBCgABBAAATowwAACqgjAAAOORMAAAzYIsAAIAAgAAAKuC3YIBAGwA5oJRAAAAFAF/AQEKAAEHAADnkDAAAOORMAAAyYIoAACqlDAAAOeQMAAA45EwAADJgigABQABAAwAv4IBADwA5oIBAJQA4oIBAMQAv4IBAPQA5oIxAAAAdADDAAEKAAEFAACzgigAALSCLgAA6oIwAADOkDAAAMyCLAABAAIAhgCigrWCMwAAAC4BIQEBCgABBwAAooIqAADtgi4AAKiCLAAAxoImAADIgi4AAOiCJAAAxIIqAAAAMwAAAM4AfwEBCgABBwAAsYIoAACvgiwAAMyCLAAA3oIsAAC3giwAANyCJgAAxYIuAAAA"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
