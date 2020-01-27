package fme

import "math"

type LyricBody struct {
	LyricCount uint16
	Lyrics     []LyricChar
	RubyCount  uint16
	Ruby       []LyricRuby
}

func (lb LyricBody) CalcByteSize() int {
	count := 0x2 + 0x2 // LyricCount, RubyCount
	for _, l := range lb.Lyrics {
		count += l.CalcByteSize()
	}
	for _, r := range lb.Ruby {
		count += r.CalcByteSize()
	}
	return count
}

func NewLyricBody(lyrics []LyricChar, ruby []LyricRuby) (*LyricBody, error) {
	// TODO: このキャストした数が上限値越えてないか見るのにもっといい方法ないかな？

	if !(0 <= len(lyrics) && len(lyrics) < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}
	if !(0 <= len(ruby) && len(ruby) < math.MaxUint16) {
		return nil, ErrBeyondBinary
	}

	return &LyricBody{uint16(len(lyrics)), lyrics, uint16(len(ruby)), ruby}, nil
}
