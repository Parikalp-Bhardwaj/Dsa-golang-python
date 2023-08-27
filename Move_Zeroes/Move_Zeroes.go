package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}

func moveZeroes(nums []int) []int {

	var idx int
	for _, num := range nums {
		nums[idx] = num
		if num != 0 {
			idx++
		}
	}

	for ; idx < len(nums); idx++ {
		fmt.Println(idx)
		nums[idx] = 0
	}

	return nums
}
