package main

import (
	"fmt"
	"sort"
)

func main() {
	satisfaction := []int{-1, -8, 0, 5, -9}
	fmt.Println(maxSatisfaction(satisfaction))
}

func maxSatisfaction(satisfaction []int) int {
	sum, ans := 0, 0
	sort.Ints(satisfaction)
	for i := len(satisfaction) - 1; i >= 0; i-- {
		if sum+satisfaction[i] > 0 {
			sum += satisfaction[i]
			ans += sum
		}
		continue
	}
	fmt.Println(ans)
	return 0
}
