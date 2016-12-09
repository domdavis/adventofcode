package day9

import "fmt"

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %s",
		len(decode(data)), decode(data))
}

func decode(input string) string {
	return input
}
