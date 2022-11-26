package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums))
}

func minimumOperations(nums []int) int {
	mk := make(map[int]int)
	for i, v := range nums {
		if v != 0 {
			mk[v] = i
		}
	}
	return len(mk)
}
