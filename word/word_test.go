package word

import (
	"testing"
)

func TestShuffleWordInnerials(t *testing.T) {

	string1 := ""
	string2 := "a"
	string3 := "it"
	string4 := "dog"

	result1 := ShuffleWordInnerials(string1)

	if result1 != string1 {
		t.Errorf("It should return same word if length is 0, got: %s, want: %s.", result1, string1)
	}

	result2 := ShuffleWordInnerials(string2)

	if result2 != string2 {
		t.Errorf("It should return same word if length is 1, got: %s, want: %s.", result2, string2)
	}

	result3 := ShuffleWordInnerials(string3)

	if result3 != string3 {
		t.Errorf("It should return same word if length is 2, got: %s, want: %s.", result3, string3)
	}

	result4 := ShuffleWordInnerials(string4)

	if result4 != string4 {
		t.Errorf("It should return same word if length is 3, got: %s, want: %s.", result4, string4)
	}
}
