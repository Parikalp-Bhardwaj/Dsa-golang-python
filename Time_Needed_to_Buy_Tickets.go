package main

import (
	"fmt"
)

func main() {
	tickets := []int{5, 1, 1, 1}
	k := 0
	fmt.Println(timeRequiredToBuy(tickets, k))
}

func timeRequiredToBuy(tickets []int, k int) int {
	t := 0
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 && tickets[k] != 0 {
				tickets[i] -= 1
				t += 1
			}
		}
	}

	return t
}
