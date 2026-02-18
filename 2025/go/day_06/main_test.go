package main

import "testing"

func TestAdd(t *testing.T) {
	values := []int{1, 2, 3, 4}
	expect := 10

	res := add(values)

	if res != expect {
		t.Errorf("Expected %v, got %v", expect, res)
	}
}

func TestMultiply(t *testing.T) {
	values := []int{1, 2, 3, 4}
	expect := 24

	res := multiply(values)

	if res != expect {
		t.Errorf("Expected %v, got %v", expect, res)
	}
}
