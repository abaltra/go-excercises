package main

import "testing"

func TestMatrixElementSum(t *testing.T) {
	input := [][]int{
		{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3},
	}

	expected := 9
	result := matrixElementsSum(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestMatrixElementSumInvalidColumn(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}

	expected := 9
	result := matrixElementsSum(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}
