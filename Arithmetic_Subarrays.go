package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))

}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	bol := []bool{}

	for i := 0; i < len(l); i++ {
		ls := []int{}
		for j := l[i]; j < r[i]+1; j++ {
			ls = append(ls, nums[j])
		}
		sort.Ints(ls)
		res := []int{}
		for k := 0; k < len(ls)-1; k++ {
			diff := math.Abs(float64(ls[k]) - float64(ls[k+1]))
			res = append(res, int(diff))

		}

		final := mapp(res)
		if len(final) == 1 {
			bol = append(bol, true)
		} else {
			bol = append(bol, false)
		}

	}

	return bol
}

func mapp(res []int) map[int]int {
	mk := make(map[int]int)
	for i, v := range res {
		mk[v] = i
	}
	return mk
}
