package main

import "testing"

func TestFindLongestSubarrayBySum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputTarget := 15
	expected := []int{1, 5}

	result := findLongestSubarrayBySum(inputTarget, input)

	if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v to match %v", result, expected)
	}
}

func TestFindLongestSubarrayBySumNotFound(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputTarget := 110000
	expected := []int{-1}

	result := findLongestSubarrayBySum(inputTarget, input)

	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("Expected %v to match %v", result, expected)
	}
}

func TestFindLongestSubarrayBySumAllZeroes(t *testing.T) {
	input := []int{0, 0, 0, 0, 0, 0, 0}
	inputTarget := 0
	expected := []int{1, 7}

	result := findLongestSubarrayBySum(inputTarget, input)

	if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v to match %v", result, expected)
	}
}

func TestFindLongestSubarrayBySumAllWrappedByZeroes(t *testing.T) {
	input := []int{0, 0, 0, 1, 0, 0, 0}
	inputTarget := 1
	expected := []int{1, 7}

	result := findLongestSubarrayBySum(inputTarget, input)

	if len(result) != len(expected) || result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v to match %v", result, expected)
	}
}
