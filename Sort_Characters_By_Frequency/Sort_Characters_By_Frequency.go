package main

import (
	"fmt"
	"sort"
)

func main() {
	// s := "tree"
	s := "Aabb"
	fmt.Println(frequencySort(s))
}

func frequencySort(s string) string {
	s1 := []byte(s)
	mk := make([]int, 128)
	for _, v := range s {
		mk[v]++
	}

	sort.Slice(s1, func(i, j int) bool {
		f1, f2 := mk[s1[i]], mk[s1[j]]
		if f1 != f2 {
			return f1 > f2
		}
		return s1[i] > s1[j]
	})
	return string(s1)
}
