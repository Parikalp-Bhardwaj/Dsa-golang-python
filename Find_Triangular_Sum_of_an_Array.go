package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(triangularSum(nums))
}

func triangularSum(nums []int) int {
	n := len(nums)
	for n > 1 {
		var ls []int
		for i := 0; i < n-1; i++ {
			ls = append(ls, (nums[i]+nums[i+1])%10)
		}
		n -= 1
		nums = ls
	}
	return nums[0]
}
