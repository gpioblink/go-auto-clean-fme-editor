package fme

type Color struct {
	RGB uint16
}

func (cl *Color) GetRGB888() (r int, g int, b int) {
	colorBin := cl.RGB
	red := (colorBin & 0b0111110000000000) >> 10
	green := (colorBin & 0b0000001111100000) >> 5
	blue := colorBin & 0b0000000000011111
	red = (red * 255) / 31
	green = (green * 255) / 31
	blue = (blue * 255) / 31
	return int(red), int(green), int(blue)
}

func NewColorFromRGB888(r int, g int, b int) *Color {
	color := uint16(((r & 0b11111000) << 7) | ((g & 0b11111000) << 2) | (b >> 3))
	return &Color{color}
}

func NewColor(color uint16) *Color {
	return &Color{color}
}

func (cl *Color) GetRGB555Binary() uint16 {
	return cl.RGB
}
