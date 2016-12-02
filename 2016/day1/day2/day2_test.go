package day2

import "fmt"

func ExampleSolution_1() {
	input :=
		`ULL
RRDDD
LURDL
UUUUD`
	fmt.Println(decode(input))
	// Output: 1985
}
