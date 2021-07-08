package main

/*
	Code signal: Almost Incerasing sequence
	Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

	Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.

	Example

	For sequence = [1, 3, 2, 1], the output should be
	almostIncreasingSequence(sequence) = false.

	There is no one element in this array that can be removed in order to get a strictly increasing sequence.
*/

func almostIncreasingSequence(sequence []int) bool {
	/*
		Runtime: O(n) While we do expand over the array twice, we only do it when we find a potentntial candidate. This means this expansion happens at most once. So while technically it's O(3n), we collapse that into O(n)

		Space: O(1) We only generate a couple ints for indexes, the rest is all in place
	*/
	for idx := 1; idx < len(sequence); idx++ {
		if sequence[idx-1] >= sequence[idx] {
			if !(isIncreasing(sequence, idx-1) || isIncreasing(sequence, idx)) {
				return false
			}
		}
	}

	return true
}

func isIncreasing(seq []int, skip int) bool {
	prev := 0
	start := 1
	stop := 0

	if skip == 0 {
		prev = 1
		start = 2
	} else if skip == len(seq)-1 {
		stop += 1
	}

	for idx := start; idx < len(seq)-stop; idx++ {
		if idx == skip {
			continue
		}

		if seq[prev] >= seq[idx] {
			return false
		}
		prev = idx
	}

	return true
}
