package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 5}
	fmt.Println(specialArray(nums))
}

func specialArray(nums []int) int {
	n := len(nums) + 1

	for i := 1; i <= n; i++ {
		var count int
		for _, j := range nums {
			if j >= i {
				count++
			}
		}
		if count == i {
			return count
		}
	}
	return -1
}
