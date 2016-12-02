package key

import "strings"

type key struct {
	Value string
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
		code = code + k.Value
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
	square := [][]*key{
		{&key{Value: "1"}, &key{Value: "2"}, &key{Value: "3"}},
		{&key{Value: "4"}, &key{Value: "5"}, &key{Value: "6"}},
		{&key{Value: "7"}, &key{Value: "8"}, &key{Value: "9"}},
	}

	arrange(square)

	SquareStart = square[1][1]
}

func arrange(k [][]*key) {
	for y := 0; y < len(k); y++ {
		for x := 0; x < len(k[y]); x++ {
			if y-1 >= 0 && k[y-1][x] != nil {
				k[y][x].up = k[y-1][x]
			}
			if y+1 < len(k) && k[y+1][x] != nil {
				k[y][x].down = k[y+1][x]
			}
			if x-1 >= 0 && k[y][x-1] != nil {
				k[y][x].left = k[y][x-1]
			}
			if x+1 < len(k[y]) && k[y][x+1] != nil {
				k[y][x].right = k[y][x+1]
			}
		}
	}
}
