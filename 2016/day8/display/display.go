package display

import (
	"regexp"
	"strconv"
)

type screen [][]bool

/*
rotate row y=0 by 4
rotate column x=0 by 1
rect 3x1
*/

var operations = map[*regexp.Regexp]func(screen, int, int){
	regexp.MustCompile(`rect ([0-9]+)x([0-9+])`):               screen.rectangle,
	regexp.MustCompile(`rotate row y=([0-9]+) by ([0-9+])`):    screen.rotateRow,
	regexp.MustCompile(`rotate column x=([0-9]+) by ([0-9+])`): screen.rotateColumn,
}

func New(x, y int) screen {
	s := make(screen, y)
	for row := range s {
		s[row] = make([]bool, x)
	}

	return s
}

func (s screen) Transform(op string) {
	for o, f := range operations {
		if o.MatchString(op) {
			r := o.FindStringSubmatch(op)
			a, _ := strconv.Atoi(r[1])
			b, _ := strconv.Atoi(r[2])
			f(s, a, b)
			return
		}
	}
}

func (s screen) Count() int {
	count := 0

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] {
				count++
			}
		}
	}

	return count
}

func (s screen) String() string {
	o := ""
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] {
				o = o + "#"
			} else {
				o = o + "."
			}
		}
		o = o + "\n"
	}

	return o
}

func (s screen) rectangle(w, h int) {
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			s[x][y] = true
		}
	}
}

func (s screen) rotateRow(r, a int) {
	row := s[r]
	right, left := row[:len(row)-a], row[len(row)-a:]
	s[r] = append(left, right...)
}

func (s screen) rotateColumn(c, a int) {
	h := len(s) - 1
	for ; a > 0; a-- {
		t := s[h][c]

		for i := h; i > 0; i-- {
			s[i][c] = s[i-1][c]
		}

		s[0][c] = t
	}
}
