package fme_test

import (
	"fmt"
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain/fme"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColor(t *testing.T) {
	testCases := []uint16{
		0x0421, 0x7fff, 0x7fe7, 0x7cbf,
		0x7e40, 0x7cbf, 0x03c0, 0x03df,
		0x00ef, 0x0140, 0x5800, 0x4411,
		0x3420, 0x0000, 0x0000,
	}

	for i, c := range testCases {
		t.Run(fmt.Sprintf("standard color %d", i), func(t *testing.T) {
			color555 := c
			threeFiveColor := fme.NewColor(color555)
			assert.EqualValues(t, color555, threeFiveColor.GetRGB555Binary())

			threeFiveR, threeFiveG, threeFiveB := threeFiveColor.GetRGB888()

			grayColorFrom888 := fme.NewColorFromRGB888(threeFiveR, threeFiveG, threeFiveB)
			assert.EqualValues(t, color555, grayColorFrom888.GetRGB555Binary())

			ReThreeFiveR, ReThreeFiveG, ReThreeFiveB := threeFiveColor.GetRGB888()
			assert.EqualValues(t, threeFiveR, ReThreeFiveR)
			assert.EqualValues(t, threeFiveG, ReThreeFiveG)
			assert.EqualValues(t, threeFiveB, ReThreeFiveB)
		})
	}
}

func TestNewColorFromRGB888(t *testing.T) {
	r := 0xff
	g := 0xff
	b := 0x39
	color := fme.NewColorFromRGB888(r, g, b)
	assert.EqualValues(t, uint16(0x7fe7), color.GetRGB555Binary())
	outR, outG, outB := color.GetRGB888()
	assert.EqualValues(t, r, outR)
	assert.EqualValues(t, g, outG)
	assert.EqualValues(t, b, outB)
}
