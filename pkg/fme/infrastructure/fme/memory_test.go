package fme_test

import (
	"encoding/base64"
	fme2 "github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/converterDomain/fme"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/infrastructure/fme"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryRepository(t *testing.T) {
	fmeData := decodeTestBytes()
	fmeStructData, err := fme2.NewFmeFromBinary(fmeData)
	assert.NoError(t, err)

	mem := fme.NewMemoryRepository()
	err = mem.Save(fmeStructData)
	assert.NoError(t, err)

	fmeSavedStructData, err := mem.Get()
	assert.NoError(t, err)
	assert.EqualValues(t, fmeStructData, fmeSavedStructData)
}

func decodeTestBytes() []byte {
	kimigayoBase64 := "Sk9ZLTAyEgAAAHcAAACvAQAAAAAcACMAKAAtADQAPQBEAE0ARgAIAAAAAAAAAIxOgqqR4wCNkYnMAIzDicwAl9GNTI7nAINMg36DS4OIAINSg2KDSgAwODQzMDM1NwCMToKqkeOCzSCQ55HjgsmUqpDnkeOCyQAhBP9/53+/fEB+v3zAA98D7wBAAQBYEUQgNAAAAAAyAAAAdAAhAQEKAAEEAABOjDAAAKqCMAAA45EwAADNgiwAAgACAAAAq4LdggEAbADmglEAAAAUAX8BAQoAAQcAAOeQMAAA45EwAADJgigAAKqUMAAA55AwAADjkTAAAMmCKAAFAAEADAC/ggEAPADmggEAlADiggEAxAC/ggEA9ADmgjEAAAB0AMMAAQoAAQUAALOCKAAAtIIuAADqgjAAAM6QMAAAzIIsAAEAAgCGAKKCtYIzAAAALgEhAQEKAAEHAACigioAAO2CLgAAqIIsAADGgiYAAMiCLgAA6IIkAADEgioAAAAzAAAAzgB/AQEKAAEHAACxgigAAK+CLAAAzIIsAADegiwAALeCLAAA3IImAADFgi4AAABAHwAAAQRcKAAAAgYCJjIAAAIAAiYyAAABHmg3AAACAQJlPAAAAgECYUEAAAIBAl5GAAACAQJbSwAAAgECV1AAAAIBCyxRAAACAQtWVwAAAgUB8FoAAAIABKRdAAACAQQiYAAAAgEEoGIAAAIBBB9lAAACAQIbagAAAgECGG8AAAIBBJZxAAACAQQUdAAAAgEEk3YAAAIBBAR5AAACBgIReQAAAgEEj3sAAAIBBA5+AAACAQnwfgAAAgEJ1YIAAAIAAwOFAAACBQFThQAAAgEDo4UAAAIGAdKHAAACAQRQigAAAgEEzowAAAIBAsuRAAACAQLHlgAAAgECxJsAAAIBAsGgAAACAQm9oQAAAgEJBasAAAIAA7mtAAACAQM3sAAAAgEEtbIAAAIBBDS1AAACAQOytwAAAgEDMLoAAAIBA668AAACAQMtvwAAAgEC6sIAAAIBAqjGAAACAQbnxwAAAgEGJskAAAIBCgjKAAACAQrnzwAAAgUC5dMAAAIAA9XWAAACAQNT2QAAAgED0dsAAAIBA1DeAAACAQJM4wAAAgECSegAAAIBAkbtAAACAQJC8gAAAgECP/cAAAIBAjv8AAACAQE4AQEAAgEBbAYBAAIBC0wHAQACAQs+DwEAAgUBUA8BAAEXUA8BAAEhYg8BAAEf"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}
