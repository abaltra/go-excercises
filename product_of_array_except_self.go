package main

/*
	https://leetcode.com/problems/product-of-array-except-self/

	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

	You must write an algorithm that runs in O(n) time and without using the division operation.



	Example 1:

	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]
*/

func productExceptSelf(nums []int) []int {
	/*
		Runtime: O(n), we go through the length of the array twice
		Memory: O(n), we create a secondary array the same size as the input
	*/

	result := make([]int, len(nums))

	for idx := 0; idx < len(result); idx++ {
		result[idx] = 1
	}

	mult := 1
	for idx := 1; idx < len(result); idx++ {
		mult *= nums[idx-1]
		result[idx] = mult
	}

	mult = 1
	for idx := len(result) - 2; idx >= 0; idx-- {
		mult *= nums[idx+1]
		result[idx] = result[idx] * mult
	}

	return result
}
