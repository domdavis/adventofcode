package day7

import (
	"fmt"
	"regexp"
	"strings"
)

var hypernet = regexp.MustCompile(`\[[a-z]*]`)

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %d",
		tlsSupportCount(data), sslSupportCount(data))
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

func sslSupportCount(input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if supportsSSL(line) {
			count++
		}
	}

	return count
}

func supportsSSL(input string) bool {
	cleaned := hypernet.ReplaceAllString(input, " ")
	matches := aba(cleaned)
	for strings.Contains(input, "[") {
		fragment := hypernet.FindString(input)
		if bab(fragment, matches) {
			return true
		}

		input = strings.Replace(input, fragment, " ", 1)
	}

	return false
}

func aba(input string) []string {
	matches := []string{}
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] && input[i] != input[i+1] {
			matches = append(matches, input[i:i+3])
		}
	}

	return matches
}

func bab(fragment string, matches []string) bool {
	for _, match := range matches {
		t := fmt.Sprintf("%c%c%c", match[1], match[0], match[1])
		if strings.Contains(fragment, t) {
			return true
		}
	}

	return false
}
