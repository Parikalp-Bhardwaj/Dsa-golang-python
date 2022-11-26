package main

import (
	"fmt"
	"sort"
)

func main() {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}
	fmt.Println(mergeSimilarItems(items1, items2))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	mk := make(map[int]int)
	// mk2 := make(map[int]int)
	for _, v := range items1 {
		mk[v[0]] += v[1]
	}

	for _, v := range items2 {
		mk[v[0]] += v[1]
	}

	out := make([][]int, 0 , len(mk))
	for i, j := range mk {
		out = append(out, []int{i, j})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i][0] < out[j][0]
	})
	return out

}
