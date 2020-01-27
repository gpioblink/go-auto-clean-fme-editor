package lyric

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Ruby struct {
	fedX       int
	rubyString string
}

func (r Ruby) FedX() int {
	return r.fedX
}

func (r Ruby) RubyString() string {
	return r.rubyString
}

func NewRuby(fedX int, rubyString string) (*Ruby, error) {
	if fedX < 0 {
		return nil, ErrInvalidLength
	}

	if _, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), rubyString); err != nil {
		return nil, ErrConvertShiftJIS
	}

	return &Ruby{fedX, rubyString}, nil
}
