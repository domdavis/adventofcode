package day8

import "fmt"

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %s",
		display(data), "Not yet done")
}

func display(input string) int {
	return len(input)
}
