package day3

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var matcher = regexp.MustCompile(`\s*([\d]+)\s+([\d]+)\s+([\d]+)`)

func Solution() string {
	return fmt.Sprintf("Part 1: %s, Part 2: %s",
		strconv.Itoa(horizontal(data)),
		strconv.Itoa(vertical(data)))
}

func horizontal(input string) int {
	return count(strings.Split(input, "\n"))
}

func vertical(input string) int {
	sides := make([]string, 3)
	total := 0
	c := 0
	for _, line := range strings.Split(input, "\n") {

		m := matcher.FindStringSubmatch(line)

		if len(m) != 4 {
			panic(fmt.Sprintf("Invalid input: '%s'", line))
		}

		for i := range sides {
			sides[i] = fmt.Sprintf("%s %s", sides[i], m[i+1])
		}

		c++

		if c == 3 {
			total += count(sides)
			sides = make([]string, 3)
			c = 0
		}

	}

	return total
}

func count(sides []string) int {
	count := 0
	for _, i := range sides {
		if triangle(i) {
			count++
		}
	}

	return count
}

func triangle(input string) bool {
	sides := make([]int, 3)
	m := matcher.FindStringSubmatch(input)

	if len(m) != 4 {
		panic(fmt.Sprintf("Invalid triangle input: '%s'", input))
	}

	for i := range sides {
		sides[i], _ = strconv.Atoi(m[i+1])
	}

	sort.Ints(sides)

	return sides[0]+sides[1] > sides[2]
}
