package fme

type LyricHeader struct {
	LyricDataSize uint16
	Flag          uint16
	X             uint16
	Y             uint16
	ColorSelectBC byte
	ColorSelectAC byte
	ColorSelectBO byte
	ColorSelectAO byte
}

func NewLyricHeaderWithStandardColorPicker(lyricBodySize int, x int, y int, beforeCharColor Color, afterCharColor Color,
	beforeOutlineColor Color, afterOutlineColor Color) (*LyricHeader, error) {

	picker := StandardColorPicker

	bc, err := picker.FindColorIndex(beforeCharColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	ac, err := picker.FindColorIndex(afterCharColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	bo, err := picker.FindColorIndex(beforeOutlineColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}
	ao, err := picker.FindColorIndex(afterOutlineColor.GetRGB555Binary())
	if err != nil {
		return nil, err
	}

	lyricDataSize := uint16(lyricBodySize + LyricHeaderSize)

	return &LyricHeader{lyricDataSize, 0x00, uint16(x), uint16(y), bc, ac, bo, ao}, nil
}
