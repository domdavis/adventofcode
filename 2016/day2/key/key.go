package key

import "strings"

type key struct {
	v     string
	up    *key
	down  *key
	left  *key
	right *key
}

var SquareStart *key
var DiamondStart *key

const Up = "U"
const Down = "D"
const Left = "L"
const Right = "R"

func Decode(start *key, input string) string {
	code := ""
	k := start

	for _, seq := range strings.Split(input, "\n") {
		k = find(k, strings.TrimSpace(seq))
		code = code + k.v
	}

	if code == "" {
		code = "Nothing found!"
	}

	return code
}

func find(k *key, code string) *key {
	var last *key
	current := k

	for i := 0; i < len(code); i++ {
		last = current
		switch string(code[i]) {
		case Up:
			current = current.up
		case Down:
			current = current.down
		case Left:
			current = current.left
		case Right:
			current = current.right
		}

		if current == nil {
			current = last
		}
	}

	return current
}

func init() {
	var emptyKeySlot *key

	square := [][]*key{
		{&key{v: "1"}, &key{v: "2"}, &key{v: "3"}},
		{&key{v: "4"}, &key{v: "5"}, &key{v: "6"}},
		{&key{v: "7"}, &key{v: "8"}, &key{v: "9"}},
	}

	diamond := [][]*key{
		{emptyKeySlot, emptyKeySlot, &key{v: "1"}, emptyKeySlot, emptyKeySlot},
		{emptyKeySlot, &key{v: "2"}, &key{v: "3"}, &key{v: "4"}, emptyKeySlot},
		{&key{v: "5"}, &key{v: "6"}, &key{v: "7"}, &key{v: "8"}, &key{v: "9"}},
		{emptyKeySlot, &key{v: "A"}, &key{v: "B"}, &key{v: "C"}, emptyKeySlot},
		{emptyKeySlot, emptyKeySlot, &key{v: "D"}, emptyKeySlot, emptyKeySlot},
	}

	arrange(square)
	arrange(diamond)

	SquareStart = square[1][1]
	DiamondStart = diamond[2][0]
}

func arrange(k [][]*key) {
	for y := 0; y < len(k); y++ {
		for x := 0; x < len(k[y]); x++ {
			if k[x][y] == nil {
				continue
			}

			if y-1 >= 0 {
				k[y][x].up = k[y-1][x]
			}
			if y+1 < len(k) {
				k[y][x].down = k[y+1][x]
			}
			if x-1 >= 0 {
				k[y][x].left = k[y][x-1]
			}
			if x+1 < len(k[y]) {
				k[y][x].right = k[y][x+1]
			}
		}
	}
}
