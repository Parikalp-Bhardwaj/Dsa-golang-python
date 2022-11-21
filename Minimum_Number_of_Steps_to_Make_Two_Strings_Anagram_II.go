package main

import (
	"fmt"
	"math"
)

func main() {
	s, t := "leetcode", "coats"
	fmt.Println(minSteps(s, t))
}

func minSteps(s, t string) int {
	mk := make(map[string]int)
	for _, s1 := range s {
		mk[string(s1)]++
	}

	for _, t1 := range t {
		mk[string(t1)]--
	}
	res := 0
	for _, i := range mk {
		res += int(math.Abs(float64(i)))
	}

	return res
}
