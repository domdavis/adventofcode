package day9

import (
	"fmt"
	"regexp"
	"strconv"
)

var marker = regexp.MustCompile(`\((\d+)x(\d+)\)`)

func Solution() string {
	return fmt.Sprintf("Part 1: %d, Part 2: %d",
		len(decode(data)), v2DecompressedLength(data))
}

func decode(input string) string {
	output := ""
	for i := 0; i < len(input); i++ {
		c := string(input[i])

		if c != "(" {
			output += c
		} else {
			length, next, repeat := getMarker(input[i:])
			i += length
			s := input[i : i+next]

			for ; repeat > 0; repeat-- {
				output += s
			}

			i += next - 1
		}
	}

	return output
}

func v2DecompressedLength(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		c := string(input[i])

		if c != "(" {
			count++
		} else {
			length, next, repeat := getMarker(input[i:])
			i += length
			s := input[i : i+next]

			for ; repeat > 0; repeat-- {
				count += v2DecompressedLength(s)
			}

			i += next - 1
		}
	}

	return count
}

func getMarker(input string) (int, int, int) {
	match := marker.FindStringSubmatch(input)
	next, _ := strconv.Atoi(match[1])
	repeat, _ := strconv.Atoi(match[2])

	return len(match[0]), next, repeat
}
