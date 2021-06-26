package main

import "testing"

func TestTwoNumberSumSortedArray(t *testing.T) {
	res := TwoNumberSum([]int{1, 2, 3}, 4)
	if len(res) != 2 {
		t.Errorf("Expected response, got empty array")
	}

	expectedResponse := []int{0, 2}

	for expected := range expectedResponse {
		found := false
		for response := range res {
			if expected == response {
				found = true
			}
		}
		if !found {
			t.Errorf("Expected %d to be in result", expected)
		}
	}
}

func TestTwoNumberSumUnsortedArray(t *testing.T) {
	res := TwoNumberSum([]int{5, 1, 3, 9, 2}, 14)
	if len(res) != 2 {
		t.Errorf("Expected response, got empty array")
	}

	expectedResponse := []int{0, 3}

	for expected := range expectedResponse {
		found := false
		for response := range res {
			if expected == response {
				found = true
			}
		}
		if !found {
			t.Errorf("Expected %d to be in result", expected)
		}
	}
}

func TestTwoNumberSumNegativeNumbers(t *testing.T) {
	res := TwoNumberSum([]int{-5, 1, -3, 9, 2}, -2)
	if len(res) != 2 {
		t.Errorf("Expected response, got empty array")
	}

	expectedResponse := []int{1, 2}

	for expected := range expectedResponse {
		found := false
		for response := range res {
			if expected == response {
				found = true
			}
		}
		if !found {
			t.Errorf("Expected %d to be in result", expected)
		}
	}
}

func TestTwoNumberSumNoResult(t *testing.T) {
	res := TwoNumberSum([]int{-5, 1, -3, 9, 2}, 20)
	if len(res) != 0 {
		t.Errorf("Expected empty array, got values")
	}
}
