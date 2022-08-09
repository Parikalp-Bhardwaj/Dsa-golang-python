package main

import (
	"fmt"
)

func main() {
	startTime := []int{1, 2, 3}
	endTime := []int{3, 2, 7}
	queryTime := 4
	fmt.Println(busyStudent(startTime, endTime, queryTime))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}
