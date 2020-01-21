package fme_test

import (
	"encoding/base64"
	informationDataPart "github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInformationDataPart_ExportBinary(t *testing.T) {
	fme := decodeInformationDataTestBytes()
	infoData, err := informationDataPart.NewInformationDataPartFromBinary(fme)
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
