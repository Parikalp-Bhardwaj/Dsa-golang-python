package main

import (
	"fmt"
	"sort"
)

func main() {
	// piles := []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	// piles := []int{2, 4, 1, 2, 7, 8}
	piles := []int{2, 4, 5}
	fmt.Println(maxCoins(piles))
}

func maxCoins(piles []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(piles)))
	fmt.Println(len(piles) / 3)
	j := 0
	sum := 0
	for i := 1; i < len(piles)/3+1; i++ {
		sum += piles[i+j]
		j += 1
	}
	
	return sum
}
