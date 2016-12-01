package position

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/domdavis/adventofcode/2016/day1/coordinate"
	"github.com/domdavis/adventofcode/2016/day1/heading"
)

var format = regexp.MustCompile("^[LR][0-9]+$")

type position struct {
	current *coordinate.Coordinate
	heading *heading.Heading
	path    map[*coordinate.Coordinate]int
}

func Start() *position {
	start := coordinate.Point(0, 0)
	path := map[*coordinate.Coordinate]int{start: 1}
	return &position{start, heading.North, path}
}

func (p *position) Move(instruction string) {
	if !format.MatchString(instruction) {
		panic(fmt.Sprintf("Invalid input: '%s'", instruction))
	}

	direction := instruction[0]
	distance, err := strconv.Atoi(instruction[1:])

	if err != nil {
		panic(fmt.Sprintf("Invalid input: '%s'", instruction))
	}

	p.heading = p.heading.Turn(direction)
	for ; distance > 0; distance-- {
		p.current = p.current.Move(p.heading.Delta())
		p.path[p.current]++
	}
}

func (p *position) Distance() int {
	return p.current.Distance()
}

func (p *position) Visited(count int) []*coordinate.Coordinate {
	locations := []*coordinate.Coordinate{}

	for k, v := range p.path {
		if v == count {
			locations = append(locations, k)
		}
	}

	return locations
}
