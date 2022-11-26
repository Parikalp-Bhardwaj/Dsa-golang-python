package main

import "fmt"

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}
	fmt.Println(rearrangeArra(nums))
}

func rearrangeArra(nums []int) []int {
	mk := make([]int, len(nums))
	p := 0
	n := 1
	for _, v := range nums {
		if v > 0 {

			mk[p] = v
			p += 2
		} else {
			mk[n] = v
			n += 2
		}
	}
	return mk
}
