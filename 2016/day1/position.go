package day1

import (
	"fmt"
	"regexp"
	"strconv"
)

var format = regexp.MustCompile("^[LR][0-9]+$")

type position struct {
	x       int
	y       int
	heading *heading
}

func start() *position {
	return &position{0, 0, north}
}

func (p *position) move(instruction string) {
	if !format.MatchString(instruction) {
		panic(fmt.Sprintf("Invalid input: '%s'", instruction))
	}

	direction := instruction[0]
	distance, err := strconv.Atoi(instruction[1:])

	if err != nil {
		panic(fmt.Sprintf("Invalid input: '%s'", instruction))
	}

	p.heading = p.heading.turn(direction)
	dx, dy := p.heading.move(distance)

	p.x += dx
	p.y += dy
}

func (p *position) distance() int {
	x := p.x
	y := p.y

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	return x + y
}
