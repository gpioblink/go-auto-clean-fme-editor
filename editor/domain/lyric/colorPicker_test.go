package lyric_test

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/editor/domain/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewColor(t *testing.T) {
	testCases := []struct {
		TestName    string
		ExpectedErr bool

		red   int
		green int
		blue  int
	}{
		{
			TestName:    "基本",
			ExpectedErr: false,
			red:         38,
			green:       72,
			blue:        76,
		},
		{
			TestName:    "上限値越えを指定",
			ExpectedErr: true,
			red:         340,
			green:       12,
			blue:        223,
		},
		{
			TestName:    "下限値越えを指定",
			ExpectedErr: true,
			red:         43,
			green:       113,
			blue:        -54,
		},
		{
			TestName:    "境界値を指定",
			ExpectedErr: false,
			red:         0,
			green:       255,
			blue:        255,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			color, err := lyric.NewColor(c.red, c.green, c.blue)
			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, c.red, color.Red())
				assert.EqualValues(t, c.green, color.Green())
				assert.EqualValues(t, c.blue, color.Blue())
			}
		})
	}
}

func TestNewColorPicker(t *testing.T) {
	testColor1, err := lyric.NewColor(23, 55, 67)
	assert.NoError(t, err)

	testColor2, err := lyric.NewColor(23, 55, 67)
	assert.NoError(t, err)

	testColor3, err := lyric.NewColor(23, 55, 67)
	assert.NoError(t, err)

	testColor4, err := lyric.NewColor(23, 55, 67)
	assert.NoError(t, err)

	colorPicker, err := lyric.NewColorPicker(*testColor1, *testColor2, *testColor3, *testColor4)
	assert.NoError(t, err)

	assert.EqualValues(t, *testColor1, colorPicker.BeforeCharColor())
	assert.EqualValues(t, *testColor2, colorPicker.AfterCharColor())
	assert.EqualValues(t, *testColor3, colorPicker.BeforeOutlineColor())
	assert.EqualValues(t, *testColor4, colorPicker.AfterOutlineColor())
}
