package day9

import "fmt"

func ExampleSolution_1() {
	fmt.Println(decode("ADVENT"))
	// Output: ADVENT
}

func ExampleSolution_2() {
	fmt.Println(decode("A(1x5)BC"))
	// Output: ABBBBBC
}

func ExampleSolution_3() {
	fmt.Println(decode("(3x3)XYZ"))
	// Output: XYZXYZXYZ
}

func ExampleSolution_4() {
	fmt.Println(decode("A(2x2)BCD(2x2)EFG"))
	// Output: ABCBCDEFEFG
}

func ExampleSolution_5() {
	fmt.Println(decode("(6x1)(1x3)A"))
	// Output: (1x3)A
}

func ExampleSolution_6() {
	fmt.Println(decode("X(8x2)(3x3)ABCY"))
	// Output: X(3x3)ABC(3x3)ABCY
}
