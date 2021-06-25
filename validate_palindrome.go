package main

import "strings"

/*
https://leetcode.com/problems/valid-palindrome/

Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases
*/

func IsValidPalindrome(s string) bool {
	/*
		Memory: O(1)
		Runtime: O(26*n) -> Could be a hard O(n) if we use a hashmap for the contains check
	*/
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1
	validChars := "abcdefghijklmnopqrstuvwxyz0123456789"

	for left < right {
		if !strings.ContainsRune(validChars, rune(s[left])) {
			left += 1
			continue
		}

		if !strings.ContainsRune(validChars, rune(s[right])) {
			right -= 1
			continue
		}

		if s[left] != s[right] {
			return false
		} else {
			left += 1
			right -= 1
		}
	}

	return true
}
