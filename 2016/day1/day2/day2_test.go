package day2

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day1/day2/key"
)

func ExampleSolution_1() {
	input :=
		`ULL
RRDDD
LURDL
UUUUD`
	fmt.Println(key.Decode(key.SquareStart, input))
	// Output: 1985
}
