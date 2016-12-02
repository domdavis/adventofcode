package key

type key struct {
	Value int
	up    *key
	down  *key
	left  *key
	right *key
}

var one = &key{Value: 1}
var two = &key{Value: 2}
var three = &key{Value: 3}
var four = &key{Value: 4}
var five = &key{Value: 5}
var six = &key{Value: 6}
var seven = &key{Value: 7}
var eight = &key{Value: 8}
var nine = &key{Value: 9}

var Start = five

const Up = "U"
const Down = "D"
const Left = "L"
const Right = "R"

func Find(k *key, code string) *key {
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
}
