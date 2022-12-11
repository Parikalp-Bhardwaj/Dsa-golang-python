package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 13, 10, 12, 31}
	fmt.Println(countDistinctIntegers(nums))
}

func countDistinctIntegers(nums []int) int {
	res := make(map[int]struct{})
	for _, num := range nums {
		res[num] = struct{}{}
		str1 := strconv.Itoa(num)
		rev, _ := strconv.Atoi(reverse(str1))
		res[int(rev)] = struct{}{}
	}


	return len(res)
}

func reverse(str1 string) (out string) {
	for _, s := range str1 {
		out = string(s) + out
	}

	return out
}
