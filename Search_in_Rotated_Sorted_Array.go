package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[start] <= nums[mid] {
			if target < nums[start] || target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if target > nums[end] || nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

	}

	return -1
}
