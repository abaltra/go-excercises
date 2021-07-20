package main

import "testing"

func TestCanFinish(t *testing.T) {
	prerequsites := [][]int{
		{1, 0},
	}

	courses := 2

	expected := true
	result := canFinish(courses, prerequsites)

	if expected != result {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}

func TestCanFinishFalse(t *testing.T) {
	prerequsites := [][]int{
		{1, 0},
		{0, 1},
	}

	courses := 2

	expected := false
	result := canFinish(courses, prerequsites)

	if expected != result {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}

func TestCanFinishMultiple(t *testing.T) {
	prerequsites := [][]int{
		{1, 4}, {2, 4}, {3, 1}, {3, 2},
	}

	courses := 5

	expected := true
	result := canFinish(courses, prerequsites)

	if expected != result {
		t.Errorf("Expected %t to equal %t", result, expected)
	}
}
