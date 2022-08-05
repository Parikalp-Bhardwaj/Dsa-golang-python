package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	// arr2 := []int{2, 1, 4, 3, 9, 6}

	arr1 := []int{28, 6, 22, 8, 44, 17}
	arr2 := []int{22, 28, 8, 6}

	fmt.Println(relativeSortArray(arr1, arr2))
}
func relativeSortArray(arr1 []int, arr2 []int) []int {

	var k int
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				arr1[k], arr1[j] = arr1[j], arr1[k]
				k++
			}
		}
	}
	sort.Ints(arr1[k:])
	return arr1
}
