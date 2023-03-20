package main

import (
	"fmt"
)

func main() {
	// n := 3
	// edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	// destination := 2
	// source := 0
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2
	n := 3
	fmt.Println(validPath(n, edges, source, destination))
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	res := make(map[int][]int)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		// fmt.Println(u, v)
		res[edge[0]] = append(res[edge[0]], v)
		res[edge[1]] = append(res[edge[1]], u)
	}
	fmt.Println(res)

	visited := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, source)
	visited[source] = true
	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]
		temp := res[val]
		fmt.Println(temp, "..")
		for i := 0; i < len(temp); i++ {
			vct := temp[i]
			if visited[vct] == false {
				visited[vct] = true
				queue = append(queue, vct)
			}
		}
		if visited[destination] {
			return true
		}
	}

	return visited[destination]
}
