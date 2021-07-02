package main

/*
	https://leetcode.com/problems/longest-palindromic-substring/submissions/

	Given a string s, return the longest palindromic substring in s.

	Example 1:

	Input: s = "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.
*/

func longestPalindrome(s string) string {
	/*
		Runtime: O(n^2) We expand around every letter in the string
		Memory: O(1) We only create a couple ints to run through the string
	*/
	left := 0
	right := 0

	for idx := 0; idx < len(s); idx++ {
		_left, _right := expandPalindrome(s, idx, idx)

		if _right-_left > right-left {
			left = _left
			right = _right
		}

		_left, _right = expandPalindrome(s, idx, idx+1)

		if _right-_left > right-left {
			left = _left
			right = _right
		}
	}

	return s[left+1 : right]

}

func expandPalindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return left, right
		}
		left--
		right++
	}

	return left, right
}
