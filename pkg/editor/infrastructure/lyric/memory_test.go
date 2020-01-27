package lyric_test

import (
	lyric_domain "github.com/gpioblink/go-auto-clean-fme-editor/pkg/editor/domain/lyric"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/editor/infrastructure/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryRepository(t *testing.T) {
	repo := lyric.NewMemoryRepository()

	assertAllProducts(t, repo, []lyric_domain.Lyric{})

	lyric1 := createLyric(t)
	addLyric(t, repo, &lyric1)

	assertAllProducts(t, repo, []lyric_domain.Lyric{lyric1})

	lyric2 := createUpdatingLyric(t)
	updateLyric(t, repo, &lyric2, 0, false)

	assertAllProducts(t, repo, []lyric_domain.Lyric{lyric2})

	addLyric(t, repo, &lyric1)
	addLyric(t, repo, &lyric1)
	addLyric(t, repo, &lyric1)
	updateLyric(t, repo, &lyric2, 2, false)

	assertAllProducts(t, repo, []lyric_domain.Lyric{lyric2, lyric1, lyric2, lyric1})

	updateLyric(t, repo, &lyric2, 1234, true)

	l := byIndex(t, repo, 0, false)
	assert.EqualValues(t, lyric2, *l)

	byIndex(t, repo, 4, true)
}

func assertAllProducts(t *testing.T, repo *lyric.MemoryRepository, expectedLyric []lyric_domain.Lyric) {
	lyrics, err := repo.ListAllLyrics()
	assert.NoError(t, err)
	assert.EqualValues(t, expectedLyric, lyrics)
}

func byIndex(t *testing.T, repo *lyric.MemoryRepository, index int, isErr bool) *lyric_domain.Lyric {
	l, err := repo.ByIndex(index)

	if isErr {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}

	return l
}

func addLyric(t *testing.T, repo *lyric.MemoryRepository, lyric *lyric_domain.Lyric) {
	err := repo.Save(lyric)
	assert.NoError(t, err)
}

func updateLyric(t *testing.T, repo *lyric.MemoryRepository, lyric *lyric_domain.Lyric, index int, isErr bool) {
	err := repo.Update(lyric, index)

	if isErr {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
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

func createUpdatingLyric(t *testing.T) lyric_domain.Lyric {
	point, err := lyric_domain.NewPoint(15, 28)
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

	char1, err := lyric_domain.NewLyricChar("き", 16, "きみ")
	assert.NoError(t, err)

	char2, err := lyric_domain.NewLyricChar("み", 23, "")
	assert.NoError(t, err)

	char3, err := lyric_domain.NewLyricChar("は", 23, "")
	assert.NoError(t, err)

	char4, err := lyric_domain.NewLyricChar("は", 23, "")
	assert.NoError(t, err)

	char5, err := lyric_domain.NewLyricChar("は", 23, "")
	assert.NoError(t, err)

	char6, err := lyric_domain.NewLyricChar("は", 23, "")
	assert.NoError(t, err)

	lyricString := []lyric_domain.LyricChar{*char1, *char2, *char3, *char4, *char5, *char6}

	testLyric, err := lyric_domain.NewLyric(*point, *colorPicker, lyricString)
	assert.NoError(t, err)

	return *testLyric
}
