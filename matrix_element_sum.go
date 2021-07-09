package main

/*
	After becoming famous, the CodeBots decided to move into a new building together. Each of the rooms has a different cost, and some of them are free, but there's a rumour that all the free rooms are haunted! Since the CodeBots are quite superstitious, they refuse to stay in any of the free rooms, or any of the rooms below any of the free rooms.

	Given matrix, a rectangular matrix of integers, where each value represents the cost of the room, your task is to return the total sum of all rooms that are suitable for the CodeBots (ie: add up all the values that don't appear below a 0).

	Example

	For

	matrix = [[0, 1, 1, 2],
			[0, 5, 0, 0],
			[2, 0, 3, 3]]
	the output should be
	matrixElementsSum(matrix) = 9.
*/

func matrixElementsSum(matrix [][]int) int {
	/*
		Runtime: O(nm) we need to go through all elements
		Memoru: O(m) since we might have to keep track of all columns. Can be lowered to O(1) if we traverse columns -> rows instead of rows -> columns
	*/
	sum := 0
	skipColumns := map[int]bool{}

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == 0 {
				skipColumns[column] = true
				continue
			}

			if _, ok := skipColumns[column]; ok {
				continue
			}

			sum += matrix[row][column]
		}
	}

	return sum
}
