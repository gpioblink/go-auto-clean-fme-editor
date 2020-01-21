package fme

type LyricDataPart struct {
	LyricHeader
	LyricBody
	LyricRuby
}

type LyricHeader struct {
	colors        [15]uint16
	lyricDataSize uint16
	flag          uint16
	x             uint16
	y             uint16
	colorSelectBC byte
	colorSelectAC byte
	colorSelectBO byte
	colorSelectAO byte
}

type LyricBody struct {
	lyricCount uint16
	lyrics     []LyricChar
}

type LyricChar struct {
	fontCode byte
	char     [2]byte
	width    uint16
}

type LyricRuby struct {
	RubyCount               uint16
	RelativeHorizontalPoint uint16
	RubyChar                []byte
}
