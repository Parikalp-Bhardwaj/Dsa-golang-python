package main

import "fmt"

func main() {
	// nums := []int{1, 3, 5, 2, 4, 8, 2, 2}
	// nums := []int{3}
	nums := []int{70, 38, 21, 22}
	fmt.Println(minMaxGame(nums))
	// fmt.Println(func1(nums))
	// func1(nums)
}

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var arr []int

	for i := 0; i < len(nums)/2; i++ {
		if i%2 == 0 {
			arr = append(arr, min(nums[2*i], nums[2*i+1]))
		} else {
			arr = append(arr, max(nums[2*i], nums[2*i+1]))
		}

	}

	return minMaxGame(arr)

}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
