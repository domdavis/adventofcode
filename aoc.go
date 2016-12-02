package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/domdavis/adventofcode/2016/day1"
	"github.com/domdavis/adventofcode/2016/day1/day2"
)

var year = flag.Int("year", 0, "The year to display")
var day = flag.Int("day", 0, "The day to display")

var solutions = map[int]map[int]func() string{
	2016: {
		1: day1.Solution,
		2: day2.Solution(),
	},
}

func main() {
	flag.Parse()

	if y, ok := solutions[*year]; ok {
		if solution, ok := y[*day]; ok {
			fmt.Println(solution())
		} else {
			usage()
		}
	} else {
		usage()
	}
}

func usage() {
	fmt.Printf("Usage: %s -year <year> -day <day>\n\n", os.Args[0])
	fmt.Println("Available options:")

	for year, v := range solutions {
		fmt.Printf("  Year %d:\n", year)

		i := 0
		days := make([]string, len(v))

		for day := range v {
			days[i] = strconv.Itoa(day)
			i++
		}

		fmt.Printf("    days: %s\n", strings.Join(days, ", "))
	}
}
