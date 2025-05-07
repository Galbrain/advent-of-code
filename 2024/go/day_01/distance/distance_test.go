package distance_test

import (
	"distance/distance"
	"fmt"
	"testing"
)

func TestGetDistance(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{0, 0, 0},
		{-1, 1, 2},
		{1, -1, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Distance(%d, %d)", test.a, test.b), func(t *testing.T) {
			res := distance.GetDistance(test.a, test.b)
			if res != test.expected {
				t.Errorf("expected %d, got %d", test.expected, res)
			}
		})
	}
}

func TestGetSimilarity(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	expected := 31

	res := distance.GetSimilarity(left, right)
	if res != expected {
		t.Fatalf("expected %d, got %d", expected, res)
	}

}
