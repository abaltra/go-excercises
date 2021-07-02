package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	input := "babad"
	expected := "bab"

	result := longestPalindrome(input)

	if expected != result {
		t.Errorf("Expected %s to equal %s", expected, result)
	}
}

func TestLongestPalindromeEvenLength(t *testing.T) {
	input := "baabad"
	expected := "baab"

	result := longestPalindrome(input)

	if expected != result {
		t.Errorf("Expected %s to equal %s", expected, result)
	}
}
