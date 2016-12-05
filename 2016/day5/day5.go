package day5

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Solution() string {
	return decode("abbhdwsy")
}

const prefix = "00000"

func decode(input string) string {
	suffix := int32(0)
	code := ""

	for i := 0; i < 8; i++ {
		searching := true

		for searching {
			data := []byte(fmt.Sprintf("%s%d", input, suffix))
			r := fmt.Sprintf("%x", md5.Sum(data))
			suffix++

			if strings.HasPrefix(r, prefix) {
				code = fmt.Sprintf("%s%s", code, string(r[5]))
				searching = false
			}
		}

	}

	return string(code)
}
