package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	// nums := []int{2, 2, 3, 1}
	// nums := []int{2, 1}
	// nums := []int{1, 2}
	fmt.Println(thirdMax(nums))
}

func thirdMax(nums []int) int {
	mk := make(map[int]int)
	for _, v := range nums {
		mk[v]++
	}

	var arr []int
	for i, _ := range mk {
		arr = append(arr, i)
	}
	sort.Ints(arr)
	if len(arr) >= 3 {

		return arr[len(arr)-3]
	} else {
		return arr[len(arr)-1]

	}

}
