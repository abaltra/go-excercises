package main

/*
	Given an array of integers, find the maximum possible sum you can get from one of its contiguous subarrays. The subarray from which this sum comes must contain at least 1 element.

	Example

	For inputArray = [-2, 2, 5, -11, 6], the output should be
	arrayMaxConsecutiveSum2(inputArray) = 7.

	The contiguous subarray that gives the maximum possible sum is [2, 5], with a sum of 7.
*/

func arrayMaxConsecutiveSum2(inputArray []int) int {
	/*
		Runtime: O(n) since we have to visit all numbers
		Memory: O(1) everything happens in place and we just create a couple ints
	*/
	bestSum := inputArray[0]
	sum := inputArray[0]
	idx := 1

	for idx < len(inputArray) {
		if inputArray[idx] > sum+inputArray[idx] {
			sum = inputArray[idx]
		} else {
			sum += inputArray[idx]
		}

		if sum > bestSum {
			bestSum = sum
		}

		idx++
	}

	return bestSum
}
