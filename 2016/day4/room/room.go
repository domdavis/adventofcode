package room

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	Real   bool
	Sector int
}

var format = regexp.MustCompile(`([a-z-]+)(\d+)\[([a-z]+)]`)

func New(code string) *room {
	m := format.FindStringSubmatch(code)

	if len(m) != 4 {
		panic(fmt.Sprintf("Invalid code: %s\n", code))
	}

	sector, err := strconv.Atoi(m[2])

	if err != nil {
		panic(fmt.Sprintf("Invalid code: %s\n", code))
	}

	return &room{Real: check(m[1], m[3]), Sector: sector}
}

func (r *room) Name() string {
	return ""
}

func check(code, checksum string) bool {
	counts := make(map[string]int)
	letters := strings.Split(strings.Join(strings.Split(code, "-"), ""), "")
	sort.Strings(letters)

	for _, l := range letters {
		counts[l]++
	}

	i := 0
	r := make(cypherText, len(counts))

	for k, v := range counts {
		r[i] = &cypher{k, v}
		i++
	}

	sort.Sort(r)
	r = r[:len(checksum)]

	for _, check := range r {
		if !strings.Contains(checksum, check.char) {
			return false
		}
		checksum = strings.Replace(checksum, check.char, "", 1)
	}

	return len(checksum) == 0
}
