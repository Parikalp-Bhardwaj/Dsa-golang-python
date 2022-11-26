package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	arr := []int{1000, 1000, 3, 7}
	fmt.Println(minSetSize(arr))
}
func minSetSize(arr []int) int {
	mk := make(map[int]int)
	for _, v := range arr {
		mk[v]++
	}

	val := make([]int, 0)
	for _, j := range mk {
		val = append(val, j)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(val)))
	out := len(arr) / 2
	res := 0
	for _, v := range val {
		out -= v
		res += 1
		if out <= 0 {
			break
		}
	}
	return res
}
