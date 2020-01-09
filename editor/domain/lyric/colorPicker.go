package lyric

import "errors"

var ErrInvalidColor = errors.New("color must be 0 to 255")

type Color struct {
	red   int
	green int
	blue  int
}

func (c Color) Red() int {
	return c.red
}

func (c Color) Green() int {
	return c.green
}

func (c Color) Blue() int {
	return c.blue
}

func NewColor(red int, green int, blue int) (*Color, error) {
	if (0 <= red && red <= 255) && (0 <= green && green <= 255) && (0 <= blue && blue <= 255) {
		return &Color{red, green, blue}, nil
	}
	return nil, ErrInvalidColor
}

type ColorPicker struct {
	beforeCharColor    Color
	afterCharColor     Color
	beforeOutlineColor Color
	afterOutlineColor  Color
}

func (cp ColorPicker) BeforeCharColor() Color {
	return cp.beforeCharColor
}

func (cp ColorPicker) AfterCharColor() Color {
	return cp.afterCharColor
}

func (cp ColorPicker) BeforeOutlineColor() Color {
	return cp.beforeOutlineColor
}

func (cp ColorPicker) AfterOutlineColor() Color {
	return cp.afterOutlineColor
}

func NewColorPicker(beforeCharColor Color, afterCharColor Color, beforeOutlineColor Color, afterOutlineColor Color) (*ColorPicker, error) {
	return &ColorPicker{beforeCharColor, afterCharColor, beforeOutlineColor, afterOutlineColor}, nil
}
