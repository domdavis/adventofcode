package day1

import (
	"fmt"
)

func ExampleDay1_1() {
	fmt.Println(move("R2, L3"))
	// Output: 5
}

func ExampleDay1_2() {
	fmt.Println(move("R2, R2, R2"))
	// Output: 2
}

func ExampleDay1_3() {
	fmt.Println(move("R5, L5, R5, R3"))
	// Output: 12
}

func ExampleDay1_4() {
	fmt.Println(locate("R8, R4, R4, R8"))
	// Output: 4
}
