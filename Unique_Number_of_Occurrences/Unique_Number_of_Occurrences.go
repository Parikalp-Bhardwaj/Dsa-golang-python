package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	m1 := make(map[int]int)

	for _, v := range arr {
		m1[v]++
	}

	m2 := make(map[int]int)
	for k, v := range m1 {
		m2[v] = k
	}

	return (len(m1) == len(m2))
}
