package safety

import (
	"fmt"
	"os"
	"slices"
	"testing"
)

type report struct {
	levels   []int
	expected bool
}

func TestReadReports(t *testing.T) {
	testFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatal("Error creating test file:", err)
	}
	defer os.Remove(testFile.Name())

	testData := "1 2 3\n3 2 1\n2 4 6\n"
	if _, err := testFile.WriteString(testData); err != nil {
		t.Fatal("Error writing data to test file:", err)
	}
	testFile.Close()

	file, err := os.Open(testFile.Name())
	if err != nil {
		t.Fatal("Something went wrong opening the testfile:", err)
	}

	reports := ReadReports(file)

	expected := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{2, 4, 6},
	}

	for i, report := range reports {
		if slices.Compare(report, expected[i]) != 0 {
			t.Fatalf("expected %v, got %v", expected[i], report)
		}
	}
}

func TestIsReportSafe(t *testing.T) {
	testReports := []report{
		{levels: []int{7, 6, 4, 2, 1}, expected: true},
		{levels: []int{1, 2, 7, 8, 9}, expected: false},
		{levels: []int{9, 7, 6, 2, 1}, expected: false},
		{levels: []int{1, 3, 2, 4, 5}, expected: false},
		{levels: []int{8, 6, 4, 4, 1}, expected: false},
		{levels: []int{1, 3, 6, 7, 9}, expected: true},
	}

	for _, report := range testReports {
		t.Run(fmt.Sprintf("Is report %v safe", report.levels), func(t *testing.T) {
			res := IsReportSafe(report.levels)
			if res != report.expected {
				t.Fatalf("expected %v, got %v", report.expected, res)
			}
		})
	}
}

func TestIsReportSafeDampened(t *testing.T) {
	testReports := []report{
		{levels: []int{7, 6, 4, 2, 1}, expected: true},
		{levels: []int{1, 2, 7, 8, 9}, expected: false},
		{levels: []int{9, 7, 6, 2, 1}, expected: false},
		{levels: []int{1, 3, 2, 4, 5}, expected: true},
		{levels: []int{8, 6, 4, 4, 1}, expected: true},
		{levels: []int{1, 3, 6, 7, 9}, expected: true},
	}

	for _, report := range testReports {
		t.Run(fmt.Sprintf("Is report %v safe", report.levels), func(t *testing.T) {
			res := IsReportSafeDampened(report.levels)
			if res != report.expected {
				t.Fatalf("expected %v, got %v", report.expected, res)
			}
		})
	}
}
