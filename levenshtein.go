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

type LevenshteinPair interface {
	Lengths() (int, int)
	Swap()
	EqualAtIndices(int, int) bool
}

func LevenshteinDistance(pair LevenshteinPair) int {
	n, m := pair.Lengths()
	// Make sure n <= m to use O(min(n,m)) space
	if n > m {
		pair.Swap()
		n, m = m, n
	}

	current := count(0, n + 1)
	var previous []int
	for i := 1; i <= m; i++ {
		previous, current = current, pad(i, n)
		for j := 1; j <= n; j++ {
			add, del := previous[j] + 1, current[j - 1] + 1
			change := previous[j - 1]
			if !pair.EqualAtIndices(j - 1, i - 1) {
				change++
			}
			current[j] = min(add, del, change)
		}
	}
	return current[n]
}

type ByLetter struct {
	a, b string
}

func (p ByLetter) Lengths() (int, int) {
	return len(p.a), len(p.b)
}

func (p ByLetter) Swap() {
	p.a, p.b = p.b, p.a
}

func (p ByLetter) EqualAtIndices(i, j int) bool {
	return p.a[i] == p.b[j]
}

type ByWord struct {
	a, b []string
}

func (p ByWord) Lengths() (int, int) {
	return len(p.a), len(p.b)
}

func (p ByWord) Swap() {
	p.a, p.b = p.b, p.a
}

func (p ByWord) EqualAtIndices(i, j int) bool {
	return p.a[i] == p.b[j]
}
