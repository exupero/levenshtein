package levenshtein

import (
	"testing"
)

func TestByLetter(t *testing.T) {
	distance := Distance(ByLetter{"hello", "hippo"})
	if distance != 3 {
		t.Errorf(`"hello" to "hippo" is %d letter edits instead of 3`, distance)
	}
}

func TestByWord(t *testing.T) {
	distance := Distance(ByWord{
		[]string{"some", "words", "are", "here"},
		[]string{"no", "words", "were", "here"},
	})
	if distance != 2 {
		t.Errorf(`"some words are here" to "no words were here" is %d word edits instead of 2`, distance)
	}
}
