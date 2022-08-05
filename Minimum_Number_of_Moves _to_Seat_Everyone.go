package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}

	fmt.Println(minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	res := 0
	for i := 0; i < len(seats); i++ {
		res += int(math.Abs(float64(seats[i]) - float64(students[i])))
	}
	return res
}
