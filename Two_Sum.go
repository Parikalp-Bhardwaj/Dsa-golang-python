package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// nums := []int{3, 3}
	// target := 6
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	mk := make(map[int]int)

	for idx, val := range nums {
		rem := target - nums[0]
		if v, ok := mk[rem]; ok {
			return []int{v, idx}
		}
		mk[val] = idx
	}

	return nil

}
