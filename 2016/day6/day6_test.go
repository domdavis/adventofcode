package day6

import "fmt"

var test = `eedadn
	drvtee
	eandsr
	raavrd
	atevrs
	tsrnev
	sdttsa
	rasrtv
	nssdts
	ntnada
	svetve
	tesnvt
	vntsnd
	vrdear
	dvrsen
	enarar`

func ExampleSolution_1() {
	fmt.Println(ecMax(test))
	// Output: easter
}

func ExampleSolution_2() {
	fmt.Println(ecMin(test))
	// Output: advent
}
