package main

import "strings"

/*
https://leetcode.com/problems/length-of-last-word/

Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.

A word is a maximal substring consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Example 2:

Input: s = " "
Output: 0
*/

func lengthOfLastWord(s string) int {
	/*
		Runtime: O(n) -> we have to go through all characters
		Space: O(m) -> where m is the number of words. This could be optimized to be O(1) using two pointers
	*/
	parts := strings.Split(s, " ")

	for right := len(parts) - 1; right >= 0; right-- {
		word := strings.Trim(parts[right], " ")
		if word != "" {
			return len(word)
		}
	}

	return 0
}
