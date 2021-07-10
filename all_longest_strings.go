package main

/*
	Given an array of strings, return another array containing all of its longest strings.

	Example

	For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
	allLongestStrings(inputArray) = ["aba", "vcd", "aba"].
*/

func allLongestStrings(inputArray []string) []string {
	/*
		Runtime: O(n) we go through all items
		Memory: O(m) where m is the length of the result
	*/

	longestLength := 0
	longestStrings := make([]string, 0)

	for _, s := range inputArray {

		if len(s) == longestLength {
			longestStrings = append(longestStrings, s)
		} else if len(s) > longestLength {
			longestLength = len(s)
			longestStrings = []string{s}
		}
	}

	return longestStrings
}
