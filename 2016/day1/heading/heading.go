package heading

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day1/coordinate"
)

type Heading struct {
	dx int
	dy int
	L  *Heading
	R  *Heading
}

var North = &Heading{dx: 0, dy: 1}
var East = &Heading{dx: 1, dy: 0}
var South = &Heading{dx: 0, dy: -1}
var West = &Heading{dx: -1, dy: 0}

var l = "L"[0]
var r = "R"[0]

func (h *Heading) Turn(d uint8) *Heading {
	switch d {
	case l:
		return h.L
	case r:
		return h.R
	default:
		panic(fmt.Sprintf("Invalid direction: '%s'", d))
	}
}

func (h *Heading) Delta() *coordinate.Coordinate {
	return coordinate.New(h.dx, h.dy)
}

func init() {
	North.L = West
	North.R = East
	East.L = North
	East.R = South
	South.L = East
	South.R = West
	West.L = South
	West.R = North
}
