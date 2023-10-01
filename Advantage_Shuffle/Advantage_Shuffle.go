package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 10, 4, 11}
	fmt.Println(advantageCount(nums1, nums2))
}

func advantageCount(nums1 []int, nums2 []int) []int32 {
	sort.Ints(nums1)
	b1 := make([][2]int, len(nums2))
	for i, n := range nums2 {
		b1[i] = [2]int{n, i}
	}

	sort.Slice(b1, func(i, j int) bool {
		return b1[i][0] < b1[j][0]
	})

	res := make([]int, len(nums1))

	left, right := 0, len(nums1)-1

	for i := len(b1) - 1; i >= 0; i-- {
		val, idx := b1[i][0], b1[i][1]
		if nums1[right] > val {
			res[idx] = nums1[right]
			right--
		} else {
			res[idx] = nums1[left]
			left++
		}
	}

	return res
}
