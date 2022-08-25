package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	pivot := 10
	fmt.Println(pivotArray(nums, pivot))
}

func pivotArray(nums []int, pivot int) []int {
	ls := []int{}
	ls2 := []int{}
	mid := []int{}
	mk := []int{}
	for _, i := range nums {
		if i < pivot {
			ls = append(ls, i)
		} else if i > pivot {
			ls2 = append(ls2, i)
		} else {
			mid = append(mid, i)
		}
	}
	mk = append(mk, ls...)
	mk = append(mk, mid...)
	mk = append(mk, ls2...)

	return mk
}
