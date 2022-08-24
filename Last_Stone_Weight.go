package main

import (
	"fmt"
	"sort"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))

}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		if stones[len(stones)-1] == stones[len(stones)-2] {
			stones = stones[:len(stones)-2]
		} else {
			stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
			stones = stones[:len(stones)-1]

			sort.Ints(stones)
		}
	}

	if len(stones) > 0 {
		return stones[0]
	}

	return 0
}
