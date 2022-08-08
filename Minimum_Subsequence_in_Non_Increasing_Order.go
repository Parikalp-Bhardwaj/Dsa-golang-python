package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 10, 9, 8}
	fmt.Println(minSubsequence(nums))

}

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	ans := []int{}
	totalSum := 0
	curVal := 0
	for _, val := range nums {
		totalSum += val
	}
	for i := len(nums) - 1; i > -1; i-- {
		curVal += nums[i]
		ans = append(ans, nums[i])

		if curVal > totalSum-curVal {
			break
		}
	}
	return ans
}
