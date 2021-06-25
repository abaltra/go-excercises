package main

import "testing"

func TestIsValidPalindrome(t *testing.T) {
	s := "anitalavalatina"
	if IsValidPalindrome(s) != true {
		t.Errorf("%s is a valid Palindrome", s)
	}
}

func TestIsInvalidPalindrome(t *testing.T) {
	s := "anitalavalatinaas"
	if IsValidPalindrome(s) != false {
		t.Errorf("%s is a not valid Palindrome", s)
	}
}

func TestIsValidPalindromeWithWhitespaces(t *testing.T) {
	s := "a l a l a    "
	if IsValidPalindrome(s) != true {
		t.Errorf("%s is a valid Palindrome", s)
	}
}
