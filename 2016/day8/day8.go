package day8

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day8/display"
)

func Solution() string {
	s := display.New(50, 6)
	s.BulkTransform(data)
	return fmt.Sprintf("Part 1: %d, Part 2:\n%s",
		s.Count(), s)
}
