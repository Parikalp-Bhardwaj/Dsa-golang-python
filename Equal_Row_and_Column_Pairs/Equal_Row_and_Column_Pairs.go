package main

import "fmt"

func main() {
	// grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}

func equalPairs(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count := 0
			for k := 0; k < len(grid); k++ {
				if grid[i][k] == grid[k][j] {
					count += 1
				} else {
					break
				}
			}
			if count == len(grid) {
				ans += 1
			}
		}

	}
	return ans
}
