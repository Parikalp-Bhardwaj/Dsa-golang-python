package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println(findCenter(edges))
}

func findCenter(edges [][]int) int {
	s1 := make(map[int]bool)
	// for i := 0; i < len(edges); i++ {
	// 	u, v := edges[0], edges[1]
	// 	fmt.Println(u, v)
	// }

	for _, i := range edges {
		u, v := i[0], i[1]
		if s1[u] {
			return u
		} else if s1[v] {
			return v
		}
		s1[u] = true
		s1[v] = true

	}

	return 0
}
