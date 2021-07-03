package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "aab"
	expected := 2
	result := lengthOfLongestSubstring(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", expected, result)
	}
}

func TestLengthOfLongestSubstringEmptyString(t *testing.T) {
	input := ""
	expected := 0
	result := lengthOfLongestSubstring(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", expected, result)
	}
}

func TestLengthOfLongestSubstringAllSameChar(t *testing.T) {
	input := "bbbb"
	expected := 1
	result := lengthOfLongestSubstring(input)

	if expected != result {
		t.Errorf("Expected %d to equal %d", expected, result)
	}
}
