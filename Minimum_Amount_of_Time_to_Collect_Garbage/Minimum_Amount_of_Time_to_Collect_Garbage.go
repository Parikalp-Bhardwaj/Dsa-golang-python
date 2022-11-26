package main

import (
	"fmt"
	"strings"
)

func main() {
	// garbage := []string{"MMM", "PGM", "GP"}
	// travel := []int{3, 10}

	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}

	fmt.Println(garbageCollection(garbage, travel))
}

func garbageCollection(garbage []string, travel []int) int {
	travel = append(travel[:1], travel...)
	travel[0] = 0
	fmt.Println(travel)
	track := []string{"G", "P", "M"}
	count := 0
	for _, s := range track {
		res := 0
		idx := 0
		for i := 0; i < len(garbage); i++ {
			c := strings.Count(garbage[i], s)
			if c > 0 {
				res += c
				idx = i
			}
		}
		sum := 0
		for j := 0; j < idx+1; j++ {
			sum += travel[j]
		}
		count += res + sum
	}
	return count
}
