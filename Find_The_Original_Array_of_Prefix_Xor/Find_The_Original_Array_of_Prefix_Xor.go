package main

import "fmt"

func main() {
	pref := []int{5, 2, 0, 3, 1}
	fmt.Println(findArray(pref))
}

func findArray(pref []int) []int {
	res := make([]int, 0, len(pref))
	res = append(res, pref[0])
	for i := 1; i < len(pref); i++ {
		res = append(res, pref[i]^pref[i-1])
	}
	return res
}
