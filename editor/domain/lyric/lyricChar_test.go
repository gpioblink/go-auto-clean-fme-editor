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

		char   string
		length int
	}{
		{
			TestName:    "基本",
			ExpectedErr: false,
			char:        "僕",
			length:      12,
		},
		{
			TestName:    "複数文字を指定",
			ExpectedErr: true,
			char:        "あいうえお",
			length:      34,
		},
		{
			TestName:    "Shift_JISで表現不可能な文字を指定",
			ExpectedErr: true,
			char:        "😢아",
			length:      2,
		},
		{
			TestName:    "歌詞がない",
			ExpectedErr: true,
			char:        "",
			length:      2,
		},
		{
			TestName:    "表示時間が負",
			ExpectedErr: true,
			char:        "",
			length:      2,
		},
		{ // TODO: 半角は実態に合わせて要挙動修正
			TestName:    "半角文字の場合",
			ExpectedErr: false,
			char:        "8",
			length:      2,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			lyricChar, err := lyric.NewLyricChar(c.char, c.length)
			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, c.char, lyricChar.Char())
				assert.EqualValues(t, c.length, lyricChar.Length())
			}
		})
	}
}
