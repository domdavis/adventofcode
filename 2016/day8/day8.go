package day8

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day8/display"
)

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %s",
		part1(data), "Not yet done")
}

func part1(input string) int {
	s := display.New(50, 6)
	s.BulkTransform(input)
	return s.Count()
}
