package main

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49

	result := maxArea(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestContainerWithMostWaterOneLevel(t *testing.T) {
	input := []int{1, 1}
	expected := 1

	result := maxArea(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestContainerWithMostWaterGated(t *testing.T) {
	input := []int{4, 3, 2, 1, 4}
	expected := 16

	result := maxArea(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}

func TestContainerWithMostWaterPeakInMiddle(t *testing.T) {
	input := []int{1, 2, 1}
	expected := 2

	result := maxArea(input)

	if result != expected {
		t.Errorf("Expected %d to equal %d", result, expected)
	}
}
