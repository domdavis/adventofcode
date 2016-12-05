package day5

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const match = "00000"
const length = 8

func Solution() string {
	return door1("abbhdwsy")
}

func door1(input string) string {
	var r string
	suffix := int32(0)
	code := ""

	for i := 0; i < length; i++ {
		r, suffix = decode(input, suffix)
		code = fmt.Sprintf("%s%s", code, string(r[len(match)]))
	}

	return code
}

func door2(input string) string {
	var r string
	suffix := int32(0)
	code := make([]string, length)

	for len(strings.Join(code, "")) != length {
		r, suffix = decode(input, suffix)
		pos, err := strconv.Atoi(string(r[len(match)]))

		if err == nil && pos < length && code[pos] == "" {
			code[pos] = string(r[len(match)+1])
		}
	}

	return strings.Join(code, "")
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
