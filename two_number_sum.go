package main

/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

import (
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	/*
		Space: O(1)
		Runtime: O(nlogn)
	*/
	sort.Ints(array)

	left := 0
	right := len(array) - 1

	for left < right {
		diff := target - (array[left] + array[right])
		if diff == 0 {
			return []int{array[left], array[right]}
		} else if diff < 0 {
			right -= 1
		} else {
			left += 1
		}
	}
	return []int{}
}
