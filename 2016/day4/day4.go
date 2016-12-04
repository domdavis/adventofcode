package day4

import (
	"strings"

	"fmt"

	"github.com/domdavis/adventofcode/2016/day4/room"
)

func Solution() string {
	rooms := real(data)
	sector := 0

	for _, r := range rooms {
		if strings.HasPrefix(r.Name(), "northpole object storage") {
			sector = r.Sector
			break
		}
	}

	return fmt.Sprintf("Part 1: %d, Part 2: %d", sum(rooms), sector)
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
