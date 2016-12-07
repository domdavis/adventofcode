package day7

import "fmt"

func ExampleSolution_1() {
	fmt.Println(abbaSupportCount("abba[mnop]qrst"))
	// Output: 1
}

func ExampleSolution_2() {
	fmt.Println(abbaSupportCount("abcd[bddb]xyyx"))
	// Output: 0
}

func ExampleSolution_3() {
	fmt.Println(abbaSupportCount("aaaa[qwer]tyui"))
	// Output: 0
}

func ExampleSolution_4() {
	fmt.Println(abbaSupportCount("ioxxoj[asdfgh]zxcvbn"))
	// Output: 1
}
