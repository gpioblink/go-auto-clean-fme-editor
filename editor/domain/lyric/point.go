package lyric

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) (*Point, error) {
	return &Point{x, y}, nil
}
