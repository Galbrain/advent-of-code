package distance_test

import (
	"distance/distance"
	"os"
	"slices"
	"testing"
)

func TestExtractPairs(t *testing.T) {
	testFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatal("Error creating temp file for testing:", err)
	}
	defer os.Remove(testFile.Name())

	testData := "1  2\n12   34\n123 456\n"
	if _, err := testFile.WriteString(testData); err != nil {
		t.Fatal("Error writing data to testfile:", err)
	}
	testFile.Close()

	file, err := os.Open(testFile.Name())
	if err != nil {
		t.Fatal("Something went wrong opening the testfile: ", err)
	}

	left, right, err := distance.ExtractPairs(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedLeft := []int{1, 12, 123}
	expectedRight := []int{2, 34, 456}

	if len(left) != len(expectedLeft) {
		t.Fatalf("expected %d left, got %d", len(left), len(expectedLeft))
	}
	if len(right) != len(expectedRight) {
		t.Fatalf("expected %d right, got %d", len(right), len(expectedRight))
	}

	resLeft := slices.Compare(left, expectedLeft)
	if resLeft != 0 {
		t.Fatalf("expected %v, got %v", expectedLeft, left)
	}

	resRight := slices.Compare(right, expectedRight)
	if resRight != 0 {
		t.Fatalf("expected %v, got %v", expectedRight, right)
	}

}
