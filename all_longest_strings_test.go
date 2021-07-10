package main

import "testing"

func TestAllLongestStrings(t *testing.T) {
	input := []string{"aba", "aa", "ad", "vcd", "aba"}
	expected := []string{"aba", "vcd", "aba"}

	result := allLongestStrings(input)

	if len(result) != len(expected) {
		t.Errorf("Expected result to have length %d, had length %d", len(result), len(expected))
	}

	for idx, e := range expected {
		if e != result[idx] {
			t.Errorf("Expected %s to match %s", result[idx], e)
		}
	}
}
