package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 3, 6}
	// k := 3
	nums := []int{0, 10}
	k := 2
	fmt.Println(maxestRangeI(nums, k))
}

func maxestRangeI(nums []int, k int) int {
	max := nums[0]
	min := max
	for _, v := range nums {
		if v > min {
			min = v
		}
		if v < max {
			max = v
		}
	}
	max += k
	min -= k
	if max >= min {
		return 0
	}
	return min - max
}
