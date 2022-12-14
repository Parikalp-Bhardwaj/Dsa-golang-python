package main

import "fmt"

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	col := make([]int, len(grid))
	row := make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			row[i] = max(row[i], grid[i][j])
			col[j] = max(col[j], grid[i][j])
		}
	}

	out := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			out += min(row[i], col[j]) - grid[i][j]
		}
	}

	return out
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
