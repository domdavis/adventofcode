package day6

import (
	"fmt"
	"strings"

	"sort"

	"github.com/domdavis/adventofcode/2016/day6/frequency"
)

func Solution() string {
	return fmt.Sprintf("Part 1: %s, Part 2: %s",
		ecMax(data), ecMin(data))
}

func ecMax(input string) string {
	msg := ""

	for _, f := range buildFrequencies(input) {
		sort.Sort(f)
		msg = msg + f[0].Char
	}

	return msg
}

func ecMin(input string) string {
	msg := ""

	for _, f := range buildFrequencies(input) {
		sort.Sort(f)
		msg = msg + f[len(f)-1].Char
	}

	return msg
}

func buildFrequencies(input string) frequency.Frequencies {
	lines := strings.Split(input, "\n")
	counts := make([]map[string]int, len(strings.TrimSpace(lines[0])))

	for i := range counts {
		counts[i] = map[string]int{}
	}

	for _, line := range lines {
		for i, char := range strings.Split(strings.TrimSpace(line), "") {
			counts[i][char]++
		}
	}

	return frequency.New(counts)
}
