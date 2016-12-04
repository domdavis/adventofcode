package day4

import (
	"strings"

	"strconv"

	"github.com/domdavis/adventofcode/2016/day4/room"
)

func Solution() string {
	rooms := real(data)
	return strconv.Itoa(sum(rooms))
}

func real(input string) room.Rooms {
	rooms := room.Rooms{}

	for _, code := range strings.Split(input, "\n") {
		r := room.New(code)

		if r.Real {
			rooms = append(rooms, r)
		}
	}

	return rooms
}

func sum(rooms room.Rooms) int {
	total := 0

	for _, r := range rooms {
		total += r.Sector
	}

	return total
}
