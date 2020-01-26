package lyric_test

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLyric(t *testing.T) {
	point, colorPicker, lyricString := createLyricContent(t)

	testLyric, err := lyric.NewLyric(point, colorPicker, lyricString)
	assert.NoError(t, err)

	assert.EqualValues(t, point, testLyric.Point())
	assert.EqualValues(t, colorPicker, testLyric.Colors())
	assert.EqualValues(t, lyricString, testLyric.Lyric())
}

func createLyricContent(t *testing.T) (lyric.Point, lyric.ColorPicker, lyric.LyricString) {
	point, err := lyric.NewPoint(3, 7)
	assert.NoError(t, err)

	color1, err := lyric.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color2, err := lyric.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color3, err := lyric.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color4, err := lyric.NewColor(112, 234, 113)
	assert.NoError(t, err)

	colorPicker, err := lyric.NewColorPicker(*color1, *color2, *color3, *color4)

	char1, err := lyric.NewLyricChar("君", 16)
	assert.NoError(t, err)

	char2, err := lyric.NewLyricChar("が", 23)
	assert.NoError(t, err)

	char3, err := lyric.NewLyricChar("与", 23)
	assert.NoError(t, err)

	char4, err := lyric.NewLyricChar("は", 23)
	assert.NoError(t, err)

	lyricString := []lyric.LyricChar{*char1, *char2, *char3, *char4}

	return *point, *colorPicker, lyricString
}
