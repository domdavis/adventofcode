package day10

import "fmt"

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %s",
		find(data, 61, 17), "Not yet done")
}

func find(input string, a, b int) int {
	return len(input) + a + b
}
