package main

/*
	Two arrays are called similar if one can be obtained from another by swapping at most one pair of elements in one of the arrays.

	Given two arrays a and b, check whether they are similar.

	Example

	For a = [1, 2, 3] and b = [1, 2, 3], the output should be
	areSimilar(a, b) = true.

	The arrays are equal, no need to swap any elements.
*/

func areSimilar(a []int, b []int) bool {
	/*
		Runtime: O(n) we go through all elements
		Memory: O(1) we only create an extra variable, the rest is in place
	*/
	if len(a) != len(b) {
		return false
	}

	idxToSwap := -1

	for idx, _ := range a {
		if a[idx] != b[idx] {
			if idxToSwap == -1 {
				idxToSwap = idx
			} else {
				a[idxToSwap], a[idx] = a[idx], a[idxToSwap]
				break
			}
		}
	}

	for idx, _ := range a {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}
