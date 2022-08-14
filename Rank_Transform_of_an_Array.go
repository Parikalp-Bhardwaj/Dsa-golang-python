package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	// arr := []int{100, 100, 100}
	arr := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr))
}

func arrayRankTransform(arr []int) []int {
	temp := append([]int{}, arr...)
	sort.Ints(temp)
	mk := make(map[int]int)
	rank := 1
	res := []int{}
	for _, v := range temp {
		if mk[v] <= 0 {
			mk[v] = rank
			rank++
		}

	}

	for _, v := range arr {
		res = append(res, mk[v])
	}

	return res
}
