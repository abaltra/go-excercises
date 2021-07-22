package main

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	input := [][]int{
		{0, 30}, {5, 10}, {15, 20},
	}

	expected := false
	result := canAttendMeetings(input)

	if expected != result {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}

func TestCanAttendMeetingsTrue(t *testing.T) {
	input := [][]int{
		{7, 10}, {15, 20},
	}

	expected := true
	result := canAttendMeetings(input)

	if expected != result {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}

func TestCanAttendMeetingsUnsorted(t *testing.T) {
	input := [][]int{
		{15, 20}, {5, 10}, {0, 30},
	}

	expected := false
	result := canAttendMeetings(input)

	if expected != result {
		t.Errorf("Expected %v to equal %v", result, expected)
	}
}
