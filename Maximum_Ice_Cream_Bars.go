package main

import (
	"fmt"
	"sort"
)

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7
	fmt.Println(maxIceCream(costs, coins))
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for _, v := range costs {
		if coins >= v {
			coins -= v
			res += 1
		}
	}
	return res
}
