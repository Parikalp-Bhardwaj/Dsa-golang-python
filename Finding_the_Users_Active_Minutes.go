package main

import "fmt"

func main() {
	logs := [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}
	k := 5
	fmt.Println(findingUsersActiveMinutes(logs, k))

}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	res := make([]int, k)
	mk := make(map[int]map[int]bool)
	for _, j := range logs {
		id, min := j[0], j[1]
		if mk[id] == nil {
			mk[id] = make(map[int]bool)
		}
		mk[id][min] = true
	}
	fmt.Println(mk)
	for _, m := range mk {
		res[len(m)-1] += 1

	}

	return res

}
