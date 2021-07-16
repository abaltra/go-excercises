package main

/*
	Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with numbers in such a way that each column, each row, and each of the nine 3 × 3 sub-grids that compose the grid all contain all of the numbers from 1 to 9 one time.

	Implement an algorithm that will check whether the given grid of numbers represents a valid Sudoku puzzle according to the layout rules described above. Note that the puzzle represented by grid does not have to be solvable.
*/

var columnMap map[string][]int
var rowMap map[string][]int

func sudoku2(grid [][]string) bool {
	/*
		Runtime: O(n^2) because we visit every number once.
		Memory: O(n^2) In a fully setup valid sudoku board, we'll save one entry per value in the grid
	*/
	columnMap = make(map[string][]int)
	rowMap = make(map[string][]int)

	for row := 0; row < len(grid); row += 3 {
		for column := 0; column < len(grid); column += 3 {
			seen := make(map[string]int, 9)
			for r := row; r < row+3; r++ {
				for c := column; c < column+3; c++ {
					if grid[r][c] == "." {
						continue
					}
					seen[grid[r][c]] += 1
					columnMap[grid[r][c]] = append(columnMap[grid[r][c]], c)
					rowMap[grid[r][c]] = append(rowMap[grid[r][c]], r)
					if seen[grid[r][c]] > 1 {
						return false
					}
				}
			}
		}
	}

	for _, rows := range rowMap {
		seen := make(map[int]int)

		for _, row := range rows {
			seen[row] += 1
			if seen[row] > 1 {
				return false
			}
		}
	}

	for _, columns := range columnMap {
		seen := make(map[int]int)

		for _, column := range columns {
			seen[column] += 1
			if seen[column] > 1 {
				return false
			}
		}
	}

	return true
}
