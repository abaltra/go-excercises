package main

import (
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
*/

func convert(s string, numRows int) string {
	/*
		Runtime: O(n*m) where n is the number of rows and m is the length of the string. This is because to build the final result we created an n*m matrix, and we fully traverse it at the end
		Memory: O(n*m)
	*/
	if numRows == 1 {
		return s
	}

	result := make([][]byte, numRows)
	for i := range result {
		result[i] = make([]byte, len(s))
	}

	rowIndex := 0
	columnIndex := 0
	letterIndex := 0

	for letterIndex < len(s) {
		result[rowIndex][columnIndex] = s[letterIndex]
		letterIndex++
		rowIndex++

		if rowIndex == numRows-1 {
			// we hit the bottom, walk it back
			for rowIndex > 0 && letterIndex < len(s) {
				result[rowIndex][columnIndex] = s[letterIndex]
				rowIndex--
				columnIndex++
				letterIndex++
			}
		}
	}

	return printMatrix(result)
}

func printMatrix(arr [][]byte) string {
	var sb strings.Builder
	for row := range arr {
		for column := range arr[row] {
			if arr[row][column] != 0 {
				sb.WriteByte(arr[row][column])
			}
		}
	}

	return sb.String()
}
