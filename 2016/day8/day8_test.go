package day8

import "fmt"

var text = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1
`

func ExampleSolotion_1() {
	fmt.Println(display(text))
}
