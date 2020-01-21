package application_test

import (
	"encoding/base64"
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateColorPalette(t *testing.T) {
	fme := decodeTestBytes()
	fmeColors, err := application.CreateColorPalette(fme, 0x77)
	assert.NoError(t, err)

	fmeExpectedColors := []application.FmeColor{
		application.NewFmeColor(0x08, 0x08, 0x08),
		application.NewFmeColor(0xff, 0xff, 0xff),
		application.NewFmeColor(0xff, 0xff, 0x39),
		application.NewFmeColor(0xff, 0x29, 0xff),
		application.NewFmeColor(0xff, 0x94, 0x00),
		application.NewFmeColor(0xff, 0x29, 0xff),
		application.NewFmeColor(0x00, 0xf6, 0x00),
		application.NewFmeColor(0x00, 0xf6, 0xff),
		application.NewFmeColor(0x00, 0x39, 0x7b),
		application.NewFmeColor(0x00, 0x52, 0x00),
		application.NewFmeColor(0xb4, 0x00, 0x00),
		application.NewFmeColor(0x8b, 0x00, 0x8b),
		application.NewFmeColor(0x6a, 0x08, 0x00),
		application.NewFmeColor(0x00, 0x00, 0x00),
		application.NewFmeColor(0x00, 0x00, 0x00),
	}

	assert.EqualValues(t, fmeExpectedColors, fmeColors)
}

func decodeTestBytes() []byte {
	kimigayoBase64 := "Sk9ZLTAyEgAAAHcAAACvAQAAAAAcACMAKAAtADQAPQBEAE0ARgAIAAAAAAAAAIxOgqqR4wCNkYnMAIzDicwAl9GNTI7nAINMg36DS4OIAINSg2KDSgAwODQzMDM1NwCMToKqkeOCzSCQ55HjgsmUqpDnkeOCyQAhBP9/53+/fEB+v3zAA98D7wBAAQBYEUQgNAAAAAAyAAAAdAAhAQEKAAEEAABOjDAAAKqCMAAA45EwAADNgiwAAgACAAAAq4LdggEAbADmglEAAAAUAX8BAQoAAQcAAOeQMAAA45EwAADJgigAAKqUMAAA55AwAADjkTAAAMmCKAAFAAEADAC/ggEAPADmggEAlADiggEAxAC/ggEA9ADmgjEAAAB0AMMAAQoAAQUAALOCKAAAtIIuAADqgjAAAM6QMAAAzIIsAAEAAgCGAKKCtYIzAAAALgEhAQEKAAEHAACigioAAO2CLgAAqIIsAADGgiYAAMiCLgAA6IIkAADEgioAAAAzAAAAzgB/AQEKAAEHAACxgigAAK+CLAAAzIIsAADegiwAALeCLAAA3IImAADFgi4AAABAHwAAAQRcKAAAAgYCJjIAAAIAAiYyAAABHmg3AAACAQJlPAAAAgECYUEAAAIBAl5GAAACAQJbSwAAAgECV1AAAAIBCyxRAAACAQtWVwAAAgUB8FoAAAIABKRdAAACAQQiYAAAAgEEoGIAAAIBBB9lAAACAQIbagAAAgECGG8AAAIBBJZxAAACAQQUdAAAAgEEk3YAAAIBBAR5AAACBgIReQAAAgEEj3sAAAIBBA5+AAACAQnwfgAAAgEJ1YIAAAIAAwOFAAACBQFThQAAAgEDo4UAAAIGAdKHAAACAQRQigAAAgEEzowAAAIBAsuRAAACAQLHlgAAAgECxJsAAAIBAsGgAAACAQm9oQAAAgEJBasAAAIAA7mtAAACAQM3sAAAAgEEtbIAAAIBBDS1AAACAQOytwAAAgEDMLoAAAIBA668AAACAQMtvwAAAgEC6sIAAAIBAqjGAAACAQbnxwAAAgEGJskAAAIBCgjKAAACAQrnzwAAAgUC5dMAAAIAA9XWAAACAQNT2QAAAgED0dsAAAIBA1DeAAACAQJM4wAAAgECSegAAAIBAkbtAAACAQJC8gAAAgECP/cAAAIBAjv8AAACAQE4AQEAAgEBbAYBAAIBC0wHAQACAQs+DwEAAgUBUA8BAAEXUA8BAAEhYg8BAAEf"
	kimigayoBytes, _ := base64.StdEncoding.DecodeString(kimigayoBase64)
	return kimigayoBytes
}