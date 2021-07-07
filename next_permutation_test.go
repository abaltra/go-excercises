package main

import "testing"

func TestNextPermutation(t *testing.T) {
	input := []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
	expected := []int{1, 5, 8, 5, 1, 3, 4, 6, 7}

	nextPermutation(input)

	for idx, val := range expected {
		if val != input[idx] {
			t.Errorf("Expected %d to equal %d", input[idx], val)
		}
	}
}

func TestNextPermutationSingleElement(t *testing.T) {
	input := []int{1}
	expected := []int{1}

	nextPermutation(input)

	for idx, val := range expected {
		if val != input[idx] {
			t.Errorf("Expected %d to equal %d", input[idx], val)
		}
	}
}

func TestNextPermutationFullReverse(t *testing.T) {
	input := []int{3, 2, 1}
	expected := []int{1, 2, 3}

	nextPermutation(input)

	for idx, val := range expected {
		if val != input[idx] {
			t.Errorf("Expected %d to equal %d", input[idx], val)
		}
	}
}
