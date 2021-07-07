package main

/*
	https://leetcode.com/problems/next-permutation/solution/

	Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

	If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

	The replacement must be in place and use only constant extra memory.



	Example 1:

	Input: nums = [1,2,3]
	Output: [1,3,2]
*/

func nextPermutation(nums []int) {
	/*
		Runtime: O(n) runtime, we will have to visit all elements in some scenarios
		Space: O(1) all operations happen in place
	*/
	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums, i+1)
}

func reverse(nums []int, left int) {
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
