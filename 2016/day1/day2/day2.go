package day2

import (
	"strconv"
	"strings"

	"github.com/domdavis/adventofcode/2016/day1/day2/key"
)

func Solution() string {
	return decode(data)
}

func decode(input string) string {
	code := ""
	k := key.Start

	for _, seq := range strings.Split(input, "\n") {
		k = key.Find(k, strings.TrimSpace(seq))
		code = code + strconv.Itoa(k.Value)
	}

	return code
}
