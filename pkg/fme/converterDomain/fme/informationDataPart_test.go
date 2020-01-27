package fme_test

import (
	"encoding/base64"
	fme2 "github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/converterDomain/fme"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInformationDataPart_ExportBinary(t *testing.T) {
	fme := decodeInformationDataTestBytes()
	infoData, err := fme2.NewInformationDataPartFromBinary(fme)
	assert.NoError(t, err)

	fmeOut, err := infoData.ExportBinary()
	assert.NoError(t, err)

	assert.EqualValues(t, fme, fmeOut)
}

func decodeInformationDataTestBytes() []byte {
	kimigayoBase64 := "AAAcACMAKAAtADQAPQBEAE0ARgAIAAAAAAAAAIxOgqqR4wCNkYnMAIzDicwAl9GNTI7nAINMg36DS4OIAINSg2KDSgAwODQzMDM1NwCMToKqkeOCzSCQ55HjgsmUqpDnkeOCyQA="
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
