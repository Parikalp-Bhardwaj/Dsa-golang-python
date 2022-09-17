package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 4, 2, 3}
	arr2 := []int{-4, -3, 6, 10, 20, 30}
	d := 3
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var count int
	var res int
	for _, i := range arr1 {
		for _, j := range arr2 {
			res = ABS(i, j)
			if res <= d {
				count++
				break
			}
		}
	}
	return len(arr1) - count

}

func ABS(i, j int) int {
	if i < j {
		return j - i
	}
	return i - j
}
