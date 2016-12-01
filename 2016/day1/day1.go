package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/domdavis/adventofcode/2016/day1/position"
)

func Solution() string {
	return fmt.Sprintf("Part 1: %s, Part 2: %s", move(data), locate(data))
}

func move(input string) string {
	p := position.Start()

	for _, i := range strings.Split(input, ", ") {
		p.Move(i)
	}

	return strconv.Itoa(p.Distance())
}

func locate(input string) string {
	p := position.Start()

	for _, i := range strings.Split(input, ", ") {
		p.Move(i)
		visited := p.Visited(2)
		if len(visited) > 0 {
			return strconv.Itoa(visited[0].Distance())
		}
	}

	return "Did not revist a location"
}
