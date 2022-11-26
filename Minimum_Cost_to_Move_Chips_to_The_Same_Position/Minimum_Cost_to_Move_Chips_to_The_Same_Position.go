package main

import "fmt"

func main() {
	position := []int{1, 2, 3}
	fmt.Println(minCostToMoveChips(position))
}

func minCostToMoveChips(position []int) int {
	even := 0
	odd := 0
	for _, pos := range position {
		if pos%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}

	if even > odd {
		return odd
	} else {
		return even
	}
}
