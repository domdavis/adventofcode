package display

import "fmt"

func ExampleNew() {
	s := New(2, 2)
	fmt.Println(s)
	// Output:
	// ..
	// ..
}

func ExampleScreen_Count() {
	s := New(10, 10)
	s.rectangle(5, 5)
	fmt.Println(s.Count())
	// Output: 25
}

func ExampleScreen_rectangle() {
	s := New(3, 3)
	s.rectangle(2, 2)
	fmt.Println(s)

	// Output:
	// ##.
	// ##.
	// ...
}

func ExampleScreen_rotateRow() {
	s := New(3, 3)
	s.rectangle(2, 2)
	s.rotateRow(0, 1)
	s.rotateRow(1, 2)
	fmt.Println(s)

	// Output:
	// .##
	// #.#
	// ...
}

func ExampleScreen_rotateColumn() {
	s := New(3, 3)
	s.rectangle(2, 2)
	s.rotateColumn(0, 1)
	s.rotateColumn(1, 2)
	fmt.Println(s)

	// Output:
	// .#.
	// #..
	// ##.
}
