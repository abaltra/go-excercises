package main

/*
	You have an unsorted array arr of non-negative integers and a number s. Find a longest contiguous subarray in arr that has a sum equal to s. Return two integers that represent its inclusive bounds. If there are several possible answers, return the one with the smallest left bound. If there are no answers, return [-1].

	Your answer should be 1-based, meaning that the first position of the array is 1 instead of 0.

	Example

	For s = 12 and arr = [1, 2, 3, 7, 5], the output should be
	findLongestSubarrayBySum(s, arr) = [2, 4].

	The sum of elements from the 2nd position to the 4th position (1-based) is equal to 12: 2 + 3 + 7.
*/

func findLongestSubarrayBySum(s int, arr []int) []int {
	/*
		Runtime: O(n^2), since we are expanding out for every pair
		Memory: O(1), it all happens in place
	*/
	bestPair := []int{-1}
	bestLength := -1
	for idx := 0; idx < len(arr); idx++ {
		localSum := 0
		for jdx := idx; jdx < len(arr); jdx++ {

			localSum += arr[jdx]

			if localSum == s {
				for ; jdx < len(arr)-1 && arr[jdx+1] == 0; jdx++ {
				}
				if jdx-idx > bestLength {
					bestPair = []int{idx + 1, jdx + 1}
					bestLength = jdx - idx
					break
				}
			}

			if localSum > s {
				break
			}

		}
	}

	return bestPair
}
