package fme

type LyricColorPicker struct { // fmeで入力される色はいつも一緒なので変数名にしちゃいました
	DarkGray         uint16
	White            uint16
	Yellow           uint16
	Pink             uint16
	Orange           uint16
	Pink2            uint16
	LightGreen       uint16
	LightBlue        uint16
	DarkBlue         uint16
	DarkGreen        uint16
	GeneralMotorsRed uint16
	Purple           uint16
	Brown            uint16
	Custom1          uint16
	Custom2          uint16
}

var StandardColorPicker = LyricColorPicker{0x0421, 0x7fff, 0x7fe7, 0x7cbf, 0x7e40,
	0x7cbf, 0x03c0, 0x03df, 0x00ef, 0x0140, 0x5800,
	0x4411, 0x3420, 0x0000, 0x0000,
}

func (cp *LyricColorPicker) FindColorIndex(rgb555 uint16) (byte, error) {
	picker := cp
	colorBin := []uint16{picker.DarkGray, picker.White, picker.Yellow, picker.Pink, picker.Orange,
		picker.Pink2, picker.LightGreen, picker.LightBlue, picker.DarkBlue, picker.DarkGreen, picker.GeneralMotorsRed,
		picker.Purple, picker.Brown, picker.Custom1, picker.Custom2}

	for i, v := range colorBin {
		if rgb555 == v {
			return byte(i), nil
		}
	}
	return 255, ErrColorNotFound
}

func (cp *LyricColorPicker) IndexToColor(index int) uint16 {
	picker := cp
	colorBin := []uint16{picker.DarkGray, picker.White, picker.Yellow, picker.Pink, picker.Orange,
		picker.Pink2, picker.LightGreen, picker.LightBlue, picker.DarkBlue, picker.DarkGreen, picker.GeneralMotorsRed,
		picker.Purple, picker.Brown, picker.Custom1, picker.Custom2}

	if len(colorBin) <= index {
		return 255
	}
	return colorBin[index]
}
