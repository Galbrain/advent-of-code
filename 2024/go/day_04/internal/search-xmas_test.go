package xmas_search

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	testfile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatal("Something went wrong creating temp file:", err)
	}

	testdata := "wer\nsdf\nxcv"
	if _, err := testfile.WriteString(testdata); err != nil {
		t.Fatal("Error writing data to testfile:", err)
	}
	testfile.Close()

	file, err := os.Open(testfile.Name())
	if err != nil {
		t.Fatal("Something went wrong opening testfile:", err)
	}

	expected := [][]string{
		{"w", "e", "r"},
		{"s", "d", "f"},
		{"x", "c", "v"},
	}

	mat := convertToSlice(file)

	for i, line := range mat {
		if slices.Compare(line, expected[i]) != 0 {
			t.Fatalf("expected %v, got %v", expected, mat)
		}
	}
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected bool
	}{
		{
			input: [][]string{
				{"w", "e", "r"},
				{"w", "e", "r"},
			},
			expected: true,
		},
		{
			input: [][]string{
				{"w", "e", "r"},
				{"w", "e"},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Check if %v is symmetric", test.input), func(t *testing.T) {
			res := isSymmetric(test.input)
			if res != test.expected {
				t.Fatalf("Expected %v, got %v", test.expected, res)
			}
		})
	}
}

func TestExtractXmas(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected int
	}{
		{
			input: [][]string{
				{"X", "M", "A", "S"},
			},
			expected: 1,
		},
		{
			input: [][]string{
				{"S", "A", "M", "X"},
			},
			expected: 1,
		},
		{
			input:    [][]string{},
			expected: 0,
		},
		{
			input: [][]string{
				{"W", "E", "R"},
			},
			expected: 0,
		},
		{
			input: [][]string{
				{"X", "A", "S", "D"},
				{"R", "M", "W", "F"},
				{"F", "D", "A", "F"},
				{"F", "E", "W", "S"},
			},
			expected: 1,
		},
		{
			input: [][]string{
				{"S", "A", "S", "D"},
				{"R", "A", "W", "F"},
				{"F", "D", "M", "F"},
				{"F", "E", "W", "X"},
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing Extract xmas for %v", test.input), func(t *testing.T) {
			res, err := extractWord(test.input)
			if err != nil {
				t.Fatal("Error extracting xmas:", err)
			}
			if res != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, res)
			}
		})
	}

}
