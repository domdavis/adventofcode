package day8

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day8/display"
)

func ExampleSolution_1() {
	test := `rect 3x2`
	s := display.New(7, 3)
	s.BulkTransform(test)
	fmt.Println(s)
	// Output:
	// ###....
	// ###....
	// .......
}

func ExampleSolution_2() {
	test := `rect 3x2
             rotate column x=1 by 1`
	s := display.New(7, 3)
	s.BulkTransform(test)
	fmt.Println(s)
	// Output:
	// #.#....
	// ###....
	// .#.....
}

func ExampleSolution_3() {
	test := `rect 3x2
             rotate column x=1 by 1
             rotate row y=0 by 4`
	s := display.New(7, 3)
	s.BulkTransform(test)
	fmt.Println(s)
	// Output:
	// ....#.#
	// ###....
	// .#.....
}

func ExampleSolution_4() {
	test := `rect 3x2
             rotate column x=1 by 1
             rotate row y=0 by 4
             rotate column x=1 by 1`
	s := display.New(7, 3)
	s.BulkTransform(test)
	fmt.Println(s)
	// Output:
	// .#..#.#
	// #.#....
	// .#.....
}
