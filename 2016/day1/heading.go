package day1

import "fmt"

type heading struct {
	dx int
	dy int
	L  *heading
	R  *heading
}

var north = &heading{dx: 0, dy: 1}
var east = &heading{dx: 1, dy: 0}
var south = &heading{dx: 0, dy: -1}
var west = &heading{dx: -1, dy: 0}

var l = "L"[0]
var r = "R"[0]

func (h *heading) turn(d uint8) *heading {
	switch d {
	case l:
		return h.L
	case r:
		return h.R
	default:
		panic(fmt.Sprintf("Invalid direction: '%s'", d))
	}
}

func (h *heading) move(d int) (int, int) {
	return h.dx * d, h.dy * d
}

func init() {
	north.L = west
	north.R = east
	east.L = north
	east.R = south
	south.L = east
	south.R = west
	west.L = south
	west.R = north
}
