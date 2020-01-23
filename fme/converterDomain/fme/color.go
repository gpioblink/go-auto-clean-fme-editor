package fme

type Color struct {
	RGB uint16
}

func (cl *Color) GetRGB888() (r uint8, g uint8, b uint8) {
	colorBin := cl.RGB
	red := (colorBin & 0b0111110000000000) >> 10
	green := (colorBin & 0b0000001111100000) >> 5
	blue := colorBin & 0b0000000000011111
	red = (red * 255) / 31
	green = (green * 255) / 31
	blue = (blue * 255) / 31
	return uint8(red), uint8(green), uint8(blue)
}

func NewColorFromRGB888(r uint8, g uint8, b uint8) *Color {
	r = (31 * r) / 255
	g = (31 * g) / 255
	b = (31 * b) / 255

	var color uint16
	color = color | (uint16(r) << 10)
	color = color | (uint16(g) << 5)
	color = color | uint16(b)

	return &Color{color}
}

func NewColor(color uint16) *Color {
	return &Color{color}
}

func (cl *Color) GetRGB555Binary() uint16 {
	return cl.RGB
}
