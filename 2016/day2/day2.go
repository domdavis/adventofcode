package day2

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day2/key"
)

func Solution() string {
	return fmt.Sprintf("Part 1: %s, Part 2: %s",
		key.Decode(key.SquareStart, data),
		key.Decode(key.DiamondStart, data))
}
