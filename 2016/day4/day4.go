package day4

import (
	"strings"

	"strconv"

	"github.com/domdavis/adventofcode/2016/day4/room"
)

func Solution() string {
	return strconv.Itoa(sum(data))
}

func sum(input string) int {
	total := 0
	for _, code := range strings.Split(input, "\n") {
		room := room.New(code)

		if room.Real {
			total += room.Sector
		}
	}

	return total
}
