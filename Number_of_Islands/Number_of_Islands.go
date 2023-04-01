package main

import "fmt"

func main() {
	// grid := [][]string{
	// 	{"1", "1", "1", "1", "0"},
	// 	{"1", "1", "0", "1", "0"},
	// 	{"1", "1", "0", "0", "0"},
	// 	{"0", "0", "0", "0", "0"}}

	grid := [][]string{{"1", "1", "1", "1", "1", "0", "1", "1", "1", "1"},
		{"1", "0", "1", "0", "1", "1", "1", "1", "1", "1"},
		{"0", "1", "1", "1", "0", "1", "1", "1", "1", "1"},
		{"1", "1", "0", "1", "1", "0", "0", "0", "0", "1"},
		{"1", "0", "1", "0", "1", "0", "0", "1", "0", "1"},
		{"1", "0", "0", "1", "1", "1", "0", "1", "0", "0"},
		{"0", "0", "1", "0", "0", "1", "1", "1", "1", "0"},
		{"1", "0", "1", "1", "1", "0", "0", "1", "1", "1"},
		{"1", "1", "1", "1", "1", "1", "1", "1", "0", "1"},
		{"1", "0", "1", "1", "1", "1", "1", "1", "1", "0"}}

	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]string) int {
	count := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "1" {
				count += 1
				numIslandsUtils(i, j, n, m, grid)
			}
		}
	}
	return count

}

func numIslandsUtils(i, j, n, m int, grid [][]string) {
	grid[i][j] = "0"
	if isValid(i+1, j, n, m, grid) {
		numIslandsUtils(i+1, j, n, m, grid)
	}
	if isValid(i, j+1, n, m, grid) {
		numIslandsUtils(i, j+1, n, m, grid)
	}
	if isValid(i-1, j, n, m, grid) {
		numIslandsUtils(i-1, j, n, m, grid)
	}
	if isValid(i, j-1, n, m, grid) {
		numIslandsUtils(i, j-1, n, m, grid)
	}

}

func isValid(i, j, n, m int, grid [][]string) bool {
	if i >= 0 && i < n && j >= 0 && j < m && grid[i][j] == "1" {
		return true
	}
	return false
}
