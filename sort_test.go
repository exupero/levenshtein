package levenshtein

import (
	"testing"
	"math/rand"
)

func TestSortByLetter(t *testing.T) {
	items := ByLetter{"hippo", "harray", "hello", "harry", "barry"}
	Sort(items)
	if len(items) != 5 {
		t.Fatalf(`Sort() did not produce %d items, only %d`, 5, len(items))
	}
	if items[0] != "hippo" {
		t.Fatalf(`Sort() did not start with first item of original array, instead "%s"`, items[0])
	}
	if items[1] != "hello" {
		t.Fatalf(`Sort() did not find "hello" as being closest to "hippo", instead "%s"`, items[1])
	}
	if items[2] != "harry" {
		t.Fatalf(`Sort() did not find "harry" as being closest to "hello", instead "%s"`, items[2])
	}
	if items[3] != "harray" {
		t.Fatalf(`Sort() did not find "harray" as being closest to "harray", instead "%s"`, items[3])
	}
	if items[4] != "barry" {
		t.Fatalf(`Sort() did not find "barry" as being closest to "harray", instead "%s"`, items[4])
	}
}

func randomStrings(n int) []string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ")
	lines := []string{}
	for i := 0; i < n; i++ {
		text := make([]rune, rand.Intn(50))
		for i := range text {
			text[i] = letters[rand.Intn(len(letters))]
		}
		lines = append(lines, string(text))
	}
	return lines
}

func TestSortArbitrary(t *testing.T) {
	// Tests that large datasets can be handled without erroring.
	items := ByLetter(randomStrings(200))
	Sort(items)
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	lines := randomStrings(300)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sort(ByLetter(lines))
	}
}
