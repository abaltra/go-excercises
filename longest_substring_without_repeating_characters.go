package main

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, find the length of the longest substring without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

func lengthOfLongestSubstring(s string) int {
	/*
		Runtime: O(n) we check all characters once
		Memory: O(1) at most we'll have a map with all the different possible charaters, might be a large number, but it depends on the size of the charater set and not the input string
	*/
	left := 0
	right := 0
	longestSubstring := 0

	seenMap := make(map[byte]int)

	for left < len(s) && right < len(s) {
		if _, ok := seenMap[s[right]]; ok {
			// we've already seen this value, check if this is the longest string and move left index
			if (len(seenMap)) > longestSubstring {
				longestSubstring = len(seenMap)
			}

			for left < right {
				if s[left] == s[right] {
					//found the one to remove
					left++
					break
				}
				delete(seenMap, s[left])
				left++
			}
		} else {
			seenMap[s[right]] = 1
		}

		right++
	}

	if len(seenMap) > longestSubstring {
		longestSubstring = len(seenMap)
	}

	return longestSubstring
}
