package key

type key struct {
	value int
	up    *key
	down  *key
	left  *key
	right *key
}

var one = &key{value: 1}
var two = &key{value: 2}
var three = &key{value: 3}
var four = &key{value: 4}
var five = &key{value: 5}
var six = &key{value: 6}
var seven = &key{value: 7}
var eight = &key{value: 8}
var nine = &key{value: 9}

var Start = five

const Up = "U"
const Down = "D"
const Left = "L"
const Right = "R"

func Find(k *key, code string) int {
	var last *key
	current := k

	for i := 0; i < len(code) && current != nil; i++ {
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
	}

	return last.value
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
	nine.right = eight
}
