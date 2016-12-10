package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/domdavis/adventofcode/2016/day10/bot"
)

var instruction = regexp.MustCompile(`^(bot|value) (\d+).* (\d+)$`)

func Parse(input string, a, b int) int {
	t := -1
	bot.Callback = func(i, j, k int) {
		if a == j && b == k {
			t = i
		}
	}

	for _, i := range strings.Split(input, "\n") {
		vars := instruction.FindStringSubmatch(i)

		if vars[1] == "bot" {
			b := bot.Bot(vars[2])
			b.Instruct(i)
		} else {
			v, _ := strconv.Atoi(vars[2])
			b := bot.Bot(vars[3])
			b.Receive(v)
		}
	}

	return t
}
