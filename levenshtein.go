package levenshtein

import (
	"strings"
)

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

func LevenshteinDistance(a, b string) int {
	n, m := len(a), len(b)
	// Make sure n <= m to use O(min(n,m)) space
	if n > m {
		a, b = b, a
		n, m = m, n
	}

	current := count(0, n + 1)
	var previous []int
	for i := 1; i <= m; i++ {
		previous, current = current, pad(i, n)
		for j := 1; j <= n; j++ {
			add, del := previous[j] + 1, current[j - 1] + 1
			change := previous[j - 1]
			if a[j - 1] != b[i - 1] {
				change++
			}
			current[j] = min(add, del, change)
		}
	}
	return current[n]
}

func LevenshteinWordDistance(a, b string) int {
	c, d := strings.Split(a, " "), strings.Split(b, " ")
	n, m := len(c), len(d)
	// Make sure n <= m to use O(min(n,m)) space
	if n > m {
		c, d = d, c
		n, m = m, n
	}

	current := count(0, n + 1)
	var previous []int
	for i := 1; i <= m; i++ {
		previous, current = current, pad(i, n)
		for j := 1; j <= n; j++ {
			add, del := previous[j] + 1, current[j - 1] + 1
			change := previous[j - 1]
			if c[j - 1] != d[i - 1] {
				change++
			}
			current[j] = min(add, del, change)
		}
	}
	return current[n]
}
