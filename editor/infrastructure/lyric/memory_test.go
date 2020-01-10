package lyric_test

import (
	lyric_domain "github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/infrastructure/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryRepository(t *testing.T) {
	repo := lyric.NewMemoryRepository()

	assertAllProducts(t, repo, []lyric_domain.Lyric{})

	lyric1 := addLyric(t, repo)

	assertAllProducts(t, repo, []lyric_domain.Lyric{*lyric1})
}

func assertAllProducts(t *testing.T, repo *lyric.MemoryRepository, expectedLyric []lyric_domain.Lyric) {
	lyrics, err := repo.ListAllLyrics()
	assert.NoError(t, err)
	assert.EqualValues(t, expectedLyric, lyrics)
}

func addLyric(t *testing.T, repo *lyric.MemoryRepository) *lyric_domain.Lyric {
	l := createLyric(t)

	err := repo.Save(&l)
	assert.NoError(t, err)

	return &l
}

func createLyric(t *testing.T) lyric_domain.Lyric {
	point, err := lyric_domain.NewPoint(3, 7)
	assert.NoError(t, err)

	color1, err := lyric_domain.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color2, err := lyric_domain.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color3, err := lyric_domain.NewColor(112, 234, 113)
	assert.NoError(t, err)

	color4, err := lyric_domain.NewColor(112, 234, 113)
	assert.NoError(t, err)

	colorPicker, err := lyric_domain.NewColorPicker(*color1, *color2, *color3, *color4)

	char1, err := lyric_domain.NewLyricChar("君", 16, "きみ")
	assert.NoError(t, err)

	char2, err := lyric_domain.NewLyricChar("が", 23, "")
	assert.NoError(t, err)

	char3, err := lyric_domain.NewLyricChar("与", 23, "")
	assert.NoError(t, err)

	char4, err := lyric_domain.NewLyricChar("は", 23, "")
	assert.NoError(t, err)

	lyricString := []lyric_domain.LyricChar{*char1, *char2, *char3, *char4}

	testLyric, err := lyric_domain.NewLyric(*point, *colorPicker, lyricString)
	assert.NoError(t, err)

	return *testLyric
}
