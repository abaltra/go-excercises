package main

import "testing"

func TestAreSimilar(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	expected := true
	result := areSimilar(a, b)

	if result != expected {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}

func TestAreSimilarSwap(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}

	expected := true
	result := areSimilar(a, b)

	if result != expected {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}

func TestAreSimilarFalse(t *testing.T) {
	a := []int{3, 1, 2}
	b := []int{1, 2, 3}

	expected := false
	result := areSimilar(a, b)

	if result != expected {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}
