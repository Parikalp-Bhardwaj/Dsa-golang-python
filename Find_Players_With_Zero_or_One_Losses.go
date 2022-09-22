package main

import (
	"fmt"
	"sort"
)

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	fmt.Println(findWinners(matches))
}

func findWinners(matches [][]int) [][]int {
	l := make(map[int]int)
	l1 := []int{}
	mk := make(map[int]int)
	w := []int{}
	for _, i := range matches {
		l[i[1]]++
		mk[i[0]] += 0
		mk[i[1]]++

	}
	

	for i, j := range mk {
		if j == 1 {
			l1 = append(l1, i)
		}
		if j == 0 {
			w = append(w, i)
		}
	}

	sort.Ints(l1)
	sort.Ints(w)

	return [][]int{w, l1}
}
