package main

import (
	"fmt"
	"math/bits"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// arr := []int{1000, 1000}

	fmt.Println(sortByBits(arr))
}

func sortByBits(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if !sortingBits(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func sortingBits(a, b int) bool {
	if bits.OnesCount(uint(a)) == bits.OnesCount(uint(b)) {
		return a < b
	}

	return bits.OnesCount(uint(a)) < bits.OnesCount(uint(b))
}
