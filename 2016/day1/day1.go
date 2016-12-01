package day1

import (
	"strconv"
	"strings"
)

func Solution() string {
	return move(data)
}

func move(input string) string {
	p := start()

	for _, i := range strings.Split(input, ", ") {
		p.move(i)
	}

	return strconv.Itoa(p.distance())
}
