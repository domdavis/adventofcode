package frequency

type frequency struct {
	Char  string
	count int
}

type Frequency []*frequency
type Frequencies []Frequency

func New(data []map[string]int) Frequencies {
	f := make(Frequencies, len(data))
	for i, char := range data {
		f[i] = Frequency{}

		for k, v := range char {
			f[i] = append(f[i], &frequency{k, v})
		}
	}

	return f
}

func (f Frequency) Len() int {
	return len(f)
}

func (f Frequency) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Frequency) Less(i, j int) bool {
	return f[i].count > f[j].count
}
