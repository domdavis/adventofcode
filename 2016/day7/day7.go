package day7

import (
	"fmt"
	"regexp"
	"strings"
)

var hypernet = regexp.MustCompile(`\[[a-z]*]`)

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %s",
		tlsSupportCount(data), "Not yet solved")
}

func tlsSupportCount(input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if supportsTLS(line) {
			count++
		}
	}

	return count
}

func supportsTLS(input string) bool {
	for strings.Contains(input, "[") {
		fragment := hypernet.FindString(input)

		if containsAbba(fragment) {
			return false
		}

		input = strings.Replace(input, fragment, " ", 1)
	}

	return containsAbba(input)
}

func containsAbba(input string) bool {
	for i := 0; i < len(input)-3; i++ {
		if abba(input[i : i+4]) {
			return true
		}
	}

	return false
}

func abba(input string) bool {
	return input[0] == input[3] && input[1] == input[2] && input[0] != input[1]
}
