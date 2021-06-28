package main

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	target := 3

	expectedLength := len(input)
	expected := []int{1, 2, 4, 5}
	removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}
}

func TestRemoveElementAllRemoved(t *testing.T) {
	input := []int{5, 5, 5, 5}
	target := 5

	expectedLength := len(input)
	expected := []int{}
	k := removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}

	if k != 0 {
		t.Errorf("Expected starting index of removed values to be 0, received %d", k)
	}
}

func TestRemoveElementRemoveStart(t *testing.T) {
	input := []int{1, 2, 3, 4}
	target := 1

	expectedLength := len(input)
	expected := []int{2, 3, 4}
	k := removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}

	if k != 3 {
		t.Errorf("Expected starting index of removed values to be 3, received %d", k)
	}
}

func TestRemoveElementRemoveEnd(t *testing.T) {
	input := []int{1, 2, 3, 4}
	target := 4

	expectedLength := len(input)
	expected := []int{1, 2, 3}
	k := removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}

	if k != 3 {
		t.Errorf("Expected starting index of removed values to be 3, received %d", k)
	}
}

func TestRemoveElementRepeated(t *testing.T) {
	input := []int{3, 2, 2, 3}
	target := 3

	expectedLength := len(input)
	expected := []int{2, 2}
	k := removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}

	if k != 2 {
		t.Errorf("Expected starting index of removed values to be 2, received %d", k)
	}
}

func TestRemoveElementNoRemoval(t *testing.T) {
	input := []int{2}
	target := 3

	expectedLength := len(input)
	expected := []int{2}
	k := removeElement(input, target)

	if len(input) != expectedLength {
		t.Errorf("Expected result array to be of length %d, was %d", expectedLength, len(input))
	}

	sub := input[:len(expected)]
	sort.Ints(sub)

	for idx, e := range expected {
		if sub[idx] != e {
			t.Errorf("Unexpected value %d in result", sub[idx])
		}
	}

	if k != 1 {
		t.Errorf("Expected starting index of removed values to be 1, received %d", k)
	}
}
