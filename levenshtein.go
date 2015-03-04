package levenshtein

func count(start, stop int) []int {
	var items []int
	for i := start; i < stop; i++ {
		items = append(items, i)
	}
	return items
}

func pad(head, tailLength int) []int {
	items := []int{head}
	for i := 0; i < tailLength; i++ {
		items = append(items, 0)
	}
	return items
}

func min(a ...int) int {
	m := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < m {
			m = i
		}
	}
	return m
}

type Sequence interface {
	Length() int
	Lengths() []int
	Swap(int, int)
	Equal(int, int) bool
	Pair(int, int) Sequence
}

func Distance(pair Sequence) int {
	lengths := pair.Lengths()
	n, m := lengths[0], lengths[1]
	// Make sure n <= m to use O(min(n,m)) space
	if n > m {
		pair.Swap(0, 1)
		n, m = m, n
	}

	current := count(0, n + 1)
	var previous []int
	for i := 1; i <= m; i++ {
		previous, current = current, pad(i, n)
		for j := 1; j <= n; j++ {
			add, del := previous[j] + 1, current[j - 1] + 1
			change := previous[j - 1]
			if !pair.Equal(j - 1, i - 1) {
				change++
			}
			current[j] = min(add, del, change)
		}
	}
	return current[n]
}

type ByLetter []string

func (p ByLetter) Length() int {
	return len(p)
}

func (p ByLetter) Lengths() []int {
	lengths := []int{}
	for _, seq := range p {
		lengths = append(lengths, len(seq))
	}
	return lengths
}

func (p ByLetter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByLetter) Equal(i, j int) bool {
	return p[0][i] == p[1][j]
}

func (p ByLetter) Pair(i, j int) Sequence {
	return ByLetter{p[i], p[j]}
}

type ByWord [][]string

func (p ByWord) Length() int {
	return len(p)
}

func (p ByWord) Lengths() []int {
	lengths := []int{}
	for _, seq := range p {
		lengths = append(lengths, len(seq))
	}
	return lengths
}

func (p ByWord) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByWord) Equal(i, j int) bool {
	return p[0][i] == p[1][j]
}

func (p ByWord) Pair(i, j int) Sequence {
	return ByWord{p[i], p[j]}
}
