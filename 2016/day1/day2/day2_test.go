package day2

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day1/day2/key"
)

const input = `ULL
RRDDD
LURDL
UUUUD`

func ExampleSolution_1() {
	fmt.Println(key.Decode(key.SquareStart, input))
	// Output: 1985
}

func ExampleSolution_2() {
	fmt.Println(key.Decode(key.DiamondStart, input))
	// Output: 5DB3
}
