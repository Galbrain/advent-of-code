package muller

import (
	"slices"
	"testing"
)

func TestGetMulMatches(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := [][]string{
		{"mul(2,4)", "2", "4"},
		{"mul(5,5)", "5", "5"},
		{"mul(11,8)", "11", "8"},
		{"mul(8,5)", "8", "5"},
	}

	matches := getMulMatches(str)
	for i, match := range matches {
		if len(match) != 3 {
			t.Fatalf("Expected match of length 3, but got %v", len(match))
		}

		if slices.Compare(match, expected[i]) != 0 {
			t.Fatalf("Expected %v, but got %v", expected[i], match)
		}
	}
}
