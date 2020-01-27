package lyric

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) (*Point, error) {
	return &Point{x, y}, nil
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}
