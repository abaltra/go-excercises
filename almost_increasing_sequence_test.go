package main

import "testing"

func TestAlmostIncreasingSequence(t *testing.T) {
	input := []int{1, 3, 2, 1}
	expected := false

	result := almostIncreasingSequence(input)

	if expected != result {
		t.Errorf("Expected result to be %t", expected)
	}
}

func TestAlmostIncreasingSequenceValid(t *testing.T) {
	input := []int{1, 3, 2}
	expected := true

	result := almostIncreasingSequence(input)

	if expected != result {
		t.Errorf("Expected result to be %t", expected)
	}
}

func TestAlmostIncreasingSequenceOnePeak(t *testing.T) {
	input := []int{1, 3, 5, 9, 1000, 10}
	expected := true

	result := almostIncreasingSequence(input)

	if expected != result {
		t.Errorf("Expected result to be %t", expected)
	}
}

func TestAlmostIncreasingSequenceAlwaysDecreasing(t *testing.T) {
	input := []int{1, 0, -1}
	expected := false

	result := almostIncreasingSequence(input)

	if expected != result {
		t.Errorf("Expected result to be %t", expected)
	}
}
