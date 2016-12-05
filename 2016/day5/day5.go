package day5

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

const match = "00000"

func Solution() string {
	return door1("abbhdwsy")
}

func door1(input string) string {
	var r string
	suffix := int32(0)
	code := ""

	for i := 0; i < 8; i++ {
		r, suffix = decode(input, suffix)
		code = fmt.Sprintf("%s%s", code, string(r[len(match)]))
	}

	return code
}

func decode(prefix string, suffix int32) (string, int32) {
	for suffix < math.MaxInt32 {
		data := []byte(fmt.Sprintf("%s%d", prefix, suffix))
		r := fmt.Sprintf("%x", md5.Sum(data))
		suffix++

		if strings.HasPrefix(r, match) {
			return r, suffix
		}
	}

	return "Found nothing", suffix
}
