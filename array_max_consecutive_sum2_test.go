package main

import "testing"

func TestArrayMaxConsecutiveSum2(t *testing.T) {
	input := []int{-2, 2, 5, -11, 6}
	expected := 7

	result := arrayMaxConsecutiveSum2(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestArrayMaxConsecutiveSum2OneElement(t *testing.T) {
	input := []int{-2}
	expected := -2

	result := arrayMaxConsecutiveSum2(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestArrayMaxConsecutiveSum2AllNegatives(t *testing.T) {
	input := []int{-2, -4, -2}
	expected := -2

	result := arrayMaxConsecutiveSum2(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}
