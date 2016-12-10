package output

var outputs = map[string]int{}

func Give(name string, v int) {
	outputs[name] = v
}

func Value(name string) int {
	return outputs[name]
}
