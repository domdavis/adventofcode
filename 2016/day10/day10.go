package day10

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day10/output"
	"github.com/domdavis/adventofcode/2016/day10/parser"
)

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %d",
		parser.Parse(data, 61, 17),
		output.Value("0")*output.Value("1")*output.Value("2"))
}
