package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

func findMaxAverage(nums []int, k int) float64 {
	var res []int
	res = append(res, nums[0:k]...)
	var sum float64
	for _, v := range res {
		sum += float64(v)
	}
	var total float64
	total = sum

	for i := 0; i < len(nums)-k; i++ {
		sum -= float64(nums[i])
		sum += float64(nums[i+k])
		if sum > total {
			total = sum
		}
	}

	return total / float64(k)
}
