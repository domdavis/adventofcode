package day4

import (
	"fmt"

	"github.com/domdavis/adventofcode/2016/day4/room"
)

func ExampleSolution_1() {
	input := "aaaaa-bbb-z-y-x-123[abxyz]"
	fmt.Println(room.New(input).Real)
	// Output: true
}

func ExampleSolution_2() {
	input := "a-b-c-d-e-f-g-h-987[abcde]"
	fmt.Println(room.New(input).Real)
	// Output: true
}

func ExampleSolution_3() {
	input := "not-a-real-room-404[oarel]"
	fmt.Println(room.New(input).Real)
	// Output: true
}

func ExampleSolution_4() {
	input := "totally-real-room-200[decoy]"
	fmt.Println(room.New(input).Real)
	// Output: false
}

func ExampleSolution_5() {
	input := `aaaaa-bbb-z-y-x-123[abxyz]
	a-b-c-d-e-f-g-h-987[abcde]
	not-a-real-room-404[oarel]
	totally-real-room-200[decoy]`
	fmt.Println(sum(input))
	// Output: 1514
}

func ExampleSolution_6() {
	input := "hqcfqwydw-fbqijys-whqii-huiuqhsx-660[qhiwf]"
	fmt.Println(room.New(input).Real)
	// Output: true
}
