package levenshtein

import (
	"testing"
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

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(ByLetter{"hippo", "harray", "hello", "harry", "barry"})
	}
}
