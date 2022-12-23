package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))

}

func findDuplicates(nums []int) []int {
	sl := make(map[int]bool)
	arr := make([]int, 0)
	for _, num := range nums {
		_, ok := sl[num]
		if ok {
			arr = append(arr, num)
		} else {
			sl[num] = true
		}
	}
	
	return arr
}
