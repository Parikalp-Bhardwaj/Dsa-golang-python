package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	fmt.Println(matrixScore(grid))
}

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] = 1 - grid[i][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if grid[j][i] == 1 {
				count += 1
			}
		}
		if count < m-count {
			for j := 0; j < m; j++ {
				grid[j][i] = 1 - grid[j][i]
			}
		}
	}
	res := 0.0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := float64(n - j - 1)
			res += float64(grid[i][j]) * math.Pow(2, p)
		}
	}
	return int(res)
}
