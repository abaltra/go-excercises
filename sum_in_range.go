package main

/*
	You have an array of integers nums and an array queries, where queries[i] is a pair of indices (0-based). Find the sum of the elements in nums from the indices at queries[i][0] to queries[i][1] (inclusive) for each query, then add all of the sums for all the queries together. Return that number modulo 109 + 7.

	Example

	For nums = [3, 0, -2, 6, -3, 2] and queries = [[0, 2], [2, 5], [0, 5]], the output should be
	sumInRange(nums, queries) = 10.

	The array of results for queries is [1, 3, 6], so the answer is 1 + 3 + 6 = 10.
*/

func sumInRange(nums []int, queries [][]int) int {
	answer := 0

	for idx := 1; idx < len(nums); idx++ {
		nums[idx] += nums[idx-1]
	}

	for _, query := range queries {
		if query[0] == 0 {
			answer += nums[query[1]]
		} else {
			answer += (nums[query[1]] - nums[query[0]-1])
		}
		answer %= 1000000007
	}

	if answer < 0 {
		answer += 1000000007
	}

	return answer
}
