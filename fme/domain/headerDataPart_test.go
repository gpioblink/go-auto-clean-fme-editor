package fme_test

import (
	"encoding/base64"
	fme "github.com/gpioblink/go-auto-clean-fme-editor/fme/domain"
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

func decodeHeaderTestBytes() []byte {
	kimigayoBase64 := "Sk9ZLTAyEgAAAHcAAACvAQAA"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
