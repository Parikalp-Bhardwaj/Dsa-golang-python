package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}
func findRelativeRanks(score []int) []string {
	s := make([]int, len(score))
	copy(s, score)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	res := make([]string, len(score))
	mk := make(map[int]string)
	for i, j := range s {
		if i == 0 {
			mk[j] = "Gold Medal"
		} else if i == 1 {
			mk[j] = "Silver Medal"
		} else if i == 2 {
			mk[j] = "Bronze Medal"
		} else {
			mk[j] = strconv.Itoa(i + 1)
		}
	}
	for j, i := range score {
		fmt.Println(i)
		res[j] = mk[i]
	}
	return res
}
