package fme_test

import (
	"encoding/base64"
	fme "github.com/gpioblink/go-auto-clean-fme-editor/fme/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimingDataPart_ExportBinary(t *testing.T) {
	fmeData := decodeTimingTestbytes()
	timingData, err := fme.NewTimingDataPartFromBinary(fmeData)
	assert.NoError(t, err)

	fmeOut, err := timingData.ExportBinary()
	assert.NoError(t, err)

	assert.EqualValues(t, fmeData, fmeOut)
}

func decodeTimingTestbytes() []byte {
	kimigayoBase64 := "QB8AAAEEXCgAAAIGAiYyAAACAAImMgAAAR5oNwAAAgECZTwAAAIBAmFBAAACAQJeRgAAAgECW0sAAAIBAldQAAACAQssUQAAAgELVlcAAAIFAfBaAAACAASkXQAAAgEEImAAAAIBBKBiAAACAQQfZQAAAgECG2oAAAIBAhhvAAACAQSWcQAAAgEEFHQAAAIBBJN2AAACAQQEeQAAAgYCEXkAAAIBBI97AAACAQQOfgAAAgEJ8H4AAAIBCdWCAAACAAMDhQAAAgUBU4UAAAIBA6OFAAACBgHShwAAAgEEUIoAAAIBBM6MAAACAQLLkQAAAgECx5YAAAIBAsSbAAACAQLBoAAAAgEJvaEAAAIBCQWrAAACAAO5rQAAAgEDN7AAAAIBBLWyAAACAQQ0tQAAAgEDsrcAAAIBAzC6AAACAQOuvAAAAgEDLb8AAAIBAurCAAACAQKoxgAAAgEG58cAAAIBBibJAAACAQoIygAAAgEK588AAAIFAuXTAAACAAPV1gAAAgEDU9kAAAIBA9HbAAACAQNQ3gAAAgECTOMAAAIBAknoAAACAQJG7QAAAgECQvIAAAIBAj/3AAACAQI7/AAAAgEBOAEBAAIBAWwGAQACAQtMBwEAAgELPg8BAAIFAVAPAQABF1APAQABIWIPAQABHw=="
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
