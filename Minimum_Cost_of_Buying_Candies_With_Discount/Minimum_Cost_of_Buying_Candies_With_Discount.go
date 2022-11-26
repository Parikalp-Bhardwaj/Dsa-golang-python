package main

import (
	"fmt"
)

func main() {
	// cost := []int{1, 2, 3}
	cost := []int{6, 5, 7, 9, 2, 2}

	// cost := []int{5, 5}
	fmt.Println(minimumCost(cost))

}

func minimumCost(cost []int) int {
	// min := cost[0]

	for i := 0; i < len(cost); i++ {
		for j := i + 1; j < len(cost); j++ {
			if cost[i] < cost[j] {
				cost[j], cost[i] = cost[i], cost[j]
			}
		}
	}

	fmt.Print(cost, "\n")

	sum := 0
	j := 0
	for _, v := range cost {
		if j == 2 {
			j = 0
			continue
		}
		sum += v
		j += 1
	}
	return sum
}
