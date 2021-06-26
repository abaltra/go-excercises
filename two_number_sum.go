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

func TwoNumberSum(array []int, target int) []int {
	/*
		Memory: O(n)
		Runtime: O(n)
	*/
	m := make(map[int]int, len(array))

	for idx, n := range array {
		if _, ok := m[n]; ok {
			return []int{m[n], idx}
		} else {
			m[target-n] = idx
		}
	}

	return []int{}
}
