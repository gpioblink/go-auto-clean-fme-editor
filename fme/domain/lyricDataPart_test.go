package fme_test

import (
	"encoding/base64"
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLyricDataPart_ExportBinary(t *testing.T) {
	fmeData := decodeLyricDataTestBytes()
	lyricData, err := fme.NewLyricDataPartFromBinary(fmeData)
	assert.NoError(t, err)

	fmeOut, err := lyricData.ExportBinary()
	assert.NoError(t, err)

	assert.EqualValues(t, fmeData, fmeOut)
}

func decodeLyricDataTestBytes() []byte {
	kimigayoBase64 := "IQT/f+d/v3xAfr98wAPfA+8AQAEAWBFEIDQAAAAAMgAAAHQAIQEBCgABBAAATowwAACqgjAAAOORMAAAzYIsAAIAAgAAAKuC3YIBAGwA5oJRAAAAFAF/AQEKAAEHAADnkDAAAOORMAAAyYIoAACqlDAAAOeQMAAA45EwAADJgigABQABAAwAv4IBADwA5oIBAJQA4oIBAMQAv4IBAPQA5oIxAAAAdADDAAEKAAEFAACzgigAALSCLgAA6oIwAADOkDAAAMyCLAABAAIAhgCigrWCMwAAAC4BIQEBCgABBwAAooIqAADtgi4AAKiCLAAAxoImAADIgi4AAOiCJAAAxIIqAAAAMwAAAM4AfwEBCgABBwAAsYIoAACvgiwAAMyCLAAA3oIsAAC3giwAANyCJgAAxYIuAAAA"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
