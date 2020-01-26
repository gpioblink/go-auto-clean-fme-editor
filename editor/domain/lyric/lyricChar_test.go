package lyric_test

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLyricChar(t *testing.T) {
	testCases := []struct {
		TestName    string
		ExpectedErr bool

		char     string
		length   int
		furigana string
	}{
		{
			TestName:    "基本",
			ExpectedErr: false,
			char:        "僕",
			length:      12,
			furigana:    "ぼく",
		},
		{
			TestName:    "ルビなし基本",
			ExpectedErr: false,
			char:        "あ",
			length:      2,
			furigana:    "",
		},
		{
			TestName:    "複数文字を指定",
			ExpectedErr: true,
			char:        "あいうえお",
			length:      34,
			furigana:    "",
		},
		{
			TestName:    "Shift_JISで表現不可能な文字を指定",
			ExpectedErr: true,
			char:        "😢아",
			length:      2,
			furigana:    "",
		},
		{
			TestName:    "Shift_JISで表現不可能なルビを指定",
			ExpectedErr: true,
			char:        "あ",
			length:      2,
			furigana:    "아아아아",
		},
		{
			TestName:    "歌詞がない",
			ExpectedErr: true,
			char:        "",
			length:      2,
			furigana:    "きみ",
		},
		{
			TestName:    "表示時間が負",
			ExpectedErr: true,
			char:        "",
			length:      2,
			furigana:    "きみ",
		},
		{ // TODO: 半角は実態に合わせて要挙動修正
			TestName:    "半角文字の場合",
			ExpectedErr: false,
			char:        "8",
			length:      2,
			furigana:    "abcd",
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			lyricChar, err := lyric.NewLyricChar(c.char, c.length, c.furigana)
			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, c.char, lyricChar.Char())
				assert.EqualValues(t, c.length, lyricChar.Length())
				assert.EqualValues(t, c.furigana, lyricChar.Furigana())
			}
		})
	}
}
