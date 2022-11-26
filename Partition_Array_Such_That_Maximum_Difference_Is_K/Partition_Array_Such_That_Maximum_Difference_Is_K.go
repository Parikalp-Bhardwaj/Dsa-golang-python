package main

import (
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 2, 5}
	k := 2
	println(partitionArray(nums, k))
}

func partitionArray(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	res := 1
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j]-nums[i] > k {
			res += 1
			j = i
		}
	}
	return res
}
