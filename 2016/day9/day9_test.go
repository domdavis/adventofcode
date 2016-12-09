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

func ExampleSolution_7() {
	fmt.Println(v2DecompressedLength("(3x3)XYZ"))
	// Output: 9
}

func ExampleSolution_8() {
	fmt.Println(v2DecompressedLength("X(8x2)(3x3)ABCY"))
	// Output: 20
}

func ExampleSolution_9() {
	fmt.Println(v2DecompressedLength("(27x12)(20x12)(13x14)(7x10)(1x12)A"))
	// Output: 241920
}

func ExampleSolution_10() {
	fmt.Println(v2DecompressedLength(
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"))
	// Output: 445
}
