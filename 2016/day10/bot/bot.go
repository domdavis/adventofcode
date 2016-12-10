package bot

import (
	"regexp"
	"strconv"
)

type bot struct {
	name        int
	values      []int
	instruction string
}

var Callback func(i, a, b int)

var bots = map[int]*bot{}
var instruction = regexp.MustCompile(`^bot.*(bot|output) (\d+).*(bot|output) (\d+)`)

func Bot(name string) *bot {
	i, _ := strconv.Atoi(name)
	if b, ok := bots[i]; ok {
		return b
	}

	b := &bot{name: i}
	bots[i] = b

	return b
}

func (b *bot) Receive(i int) {
	b.values = append(b.values, i)
	b.process()
}

func (b *bot) Instruct(i string) {
	b.instruction = i
	b.process()
}

func (b *bot) High() int {
	return b.value(func() bool { return b.values[0] >= b.values[1] })
}

func (b *bot) Low() int {
	return b.value(func() bool { return b.values[0] <= b.values[1] })
}

func (b *bot) value(f func() bool) int {
	var v int

	switch {
	case len(b.values) == 1:
		v = b.values[0]
	case f():
		v = b.values[0]
	default:
		v = b.values[1]
	}

	return v
}

func (b *bot) process() {

	if len(b.values) == 2 {
		Callback(b.name, b.values[0], b.values[1])
	}

	if len(b.values) == 2 && b.instruction != "" {
		vars := instruction.FindStringSubmatch(b.instruction)
		l := b.Low()
		h := b.High()
		b.values = []int{}
		b.give(vars[1], vars[2], l)
		b.give(vars[3], vars[4], h)
	}
}

func (b *bot) give(r, i string, v int) {
	if r == "bot" {
		Bot(i).Receive(v)
	}
}
