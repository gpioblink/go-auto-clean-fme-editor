package fme_test

import (
	"encoding/base64"
	fme "github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaderDataPart_ExportBinary(t *testing.T) {
	fmeData := decodeHeaderTestBytes()
	headData, err := fme.NewHeaderDataPartFromBinary(fmeData)
	assert.NoError(t, err)

	fmeOut, err := headData.ExportBinary()
	assert.NoError(t, err)

	assert.EqualValues(t, fmeData, fmeOut)
}

func TestCheckMagicValue(t *testing.T) {
	fmeData := decodeHeaderTestBytes()
	err := fme.CheckMagicValue(fmeData)
	assert.NoError(t, err)
}

func TestGetOffset(t *testing.T) {
	fmeData := decodeHeaderTestBytes()
	info, lyric, timing, err := fme.GetOffsets(fmeData)
	assert.NoError(t, err)
	assert.EqualValues(t, 0x12, info)
	assert.EqualValues(t, 0x77, lyric)
	assert.EqualValues(t, 0x1af, timing)
}

func decodeHeaderTestBytes() []byte {
	kimigayoBase64 := "Sk9ZLTAyEgAAAHcAAACvAQAA"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
