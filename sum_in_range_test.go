package main

import "testing"

func TestSumInRange(t *testing.T) {
	input := []int{3, 0, -2, 6, -3, 2}
	queries := [][]int{
		{0, 2},
		{2, 5},
		{0, 5},
	}

	expected := 10
	result := sumInRange(input, queries)

	if expected != result {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestSumInRangeNegative(t *testing.T) {
	input := []int{-1000}
	queries := [][]int{
		{0, 0},
	}

	expected := 999999007
	result := sumInRange(input, queries)

	if expected != result {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}
