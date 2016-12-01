package coordinate

import "fmt"

type Coordinate struct {
	x int
	y int
}

const key = "%s.%s"

var points = map[string]*Coordinate{}

func New(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

func Point(x, y int) *Coordinate {
	k := fmt.Sprintf(key, x, y)

	if p, ok := points[k]; ok {
		return p
	}

	p := New(x, y)
	points[k] = p
	return p
}

func Visited(c *Coordinate) bool {
	_, visited := points[fmt.Sprintf(key, c.x, c.y)]
	return visited
}

func (c *Coordinate) Move(d *Coordinate) *Coordinate {
	return Point(c.x+d.x, c.y+d.y)
}

func (c *Coordinate) Distance() int {
	return abs(c.x) + abs(c.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
