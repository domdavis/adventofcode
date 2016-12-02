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
	init_square()
	init_diamond()
}

func init_square() {
	one := &key{Value: "1"}
	two := &key{Value: "2"}
	three := &key{Value: "3"}
	four := &key{Value: "4"}
	five := &key{Value: "5"}
	six := &key{Value: "6"}
	seven := &key{Value: "7"}
	eight := &key{Value: "8"}
	nine := &key{Value: "9"}

	one.right = two
	one.down = four
	two.left = one
	two.right = three
	two.down = five
	three.left = two
	three.down = six
	four.up = one
	four.right = five
	four.down = seven
	five.up = two
	five.right = six
	five.down = eight
	five.left = four
	six.up = three
	six.down = nine
	six.left = five
	seven.up = four
	seven.right = eight
	eight.up = five
	eight.right = nine
	eight.left = seven
	nine.up = six
	nine.left = eight

	SquareStart = five
}

func init_diamond() {

}
