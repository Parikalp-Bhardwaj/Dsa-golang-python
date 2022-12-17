package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	fmt.Println(diagonalSort(mat))
}

func diagonalSort(mat [][]int) [][]int {
	hm := make([][]int, len(mat)*len(mat[0]))
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := j - i
			if k >= 0 {
				hm[k] = append(hm[k], mat[i][j])
			} else {
				k = k * -n
				hm[k] = append(hm[k], mat[i][j])
			}
		}
	}
	for i := range hm {
		// sort.Slice(hm[i], func(x, y int) bool {
		// 	return hm[i][x] < hm[i][y]
		// })
		sort.Sort(sort.IntSlice(hm[i]))

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := j - i
			if k >= 0 {
				mat[i][j] = hm[k][0]
				hm[k] = hm[k][1:]
			} else {
				k *= -n
				mat[i][j] = hm[k][0]
				hm[k] = hm[k][1:]
			}
		}
	}
	return mat
}
