package day4

import "fmt"

func ExampleSolution_1() {
	input := "aaaaa-bbb-z-y-x-123[abxyz]"
	fmt.Println(input)
	// Output: true
}

func ExampleSolution_2() {
	input := "a-b-c-d-e-f-g-h-987[abcde]"
	fmt.Println(input)
	// Output: true
}

func ExampleSolution_3() {
	input := "not-a-real-room-404[oarel]"
	fmt.Println(input)
	// Output: true
}

func ExampleSolution_4() {
	input := "totally-real-room-200[decoy]"
	fmt.Println(input)
	// Output: false
}

func ExampleSolution_5() {
	input := `aaaaa-bbb-z-y-x-123[abxyz]
	a-b-c-d-e-f-g-h-987[abcde]
	not-a-real-room-404[oarel]
	totally-real-room-200[decoy]`
	fmt.Println(input)
	// Output: 1514
}
