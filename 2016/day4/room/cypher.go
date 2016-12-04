package room

type cypher struct {
	char  string
	count int
}

type cypherText []*cypher

func (c cypherText) Len() int {
	return len(c)
}

func (c cypherText) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c cypherText) Less(i, j int) bool {
	if c[i].count == 1 && c[j].count == 1 {
		return c[i].char < c[j].char
	}

	return c[i].count > c[j].count
}
