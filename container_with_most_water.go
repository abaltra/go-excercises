package main

/*
	https://leetcode.com/problems/container-with-most-water

	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
	Notice that you may not slant the container.

	Input: height = [1,8,6,2,5,4,8,3,7]
	Output: 49
	Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

*/
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxArea(height []int) int {
	/*
		Runtime: O(n) as we'll go through all values in the array
		Memory: O(1), we'll only create a couple of ints to keep track of indexes and total area
	*/
	left := 0
	right := len(height) - 1
	currentBest := 0

	for left < right {
		area := (right - left) * (min(height[left], height[right]))
		if area > currentBest {
			currentBest = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return currentBest
}
