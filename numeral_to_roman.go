package main

import "strings"

/*
	https://leetcode.com/problems/integer-to-roman/

	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

	Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.
	Given an integer, convert it to a roman numeral.
*/

func intToRoman(num int) string {
	/*
		Runtime: O(n) over the times we'll have to substract elements from the base number
		Memory: O(1) we create a static map with the valid values and substract from the input. You could make an argument for O(m) for the output string
	*/
	result := strings.Builder{}
	type Pair struct {
		Key   string
		Value int
	}

	baseNumbers := []Pair{
		{Key: "M", Value: 1000},
		{Key: "CM", Value: 900},
		{Key: "D", Value: 500},
		{Key: "CD", Value: 400},
		{Key: "C", Value: 100},
		{Key: "XC", Value: 90},
		{Key: "L", Value: 50},
		{Key: "XL", Value: 40},
		{Key: "X", Value: 10},
		{Key: "IX", Value: 9},
		{Key: "V", Value: 5},
		{Key: "IV", Value: 4},
		{Key: "I", Value: 1},
	}

	idx := 0
	for num > 0 && idx < len(baseNumbers) {
		if baseNumbers[idx].Value > num {
			idx += 1
			continue
		}

		result.WriteString(baseNumbers[idx].Key)
		num -= baseNumbers[idx].Value
	}

	return result.String()
}
