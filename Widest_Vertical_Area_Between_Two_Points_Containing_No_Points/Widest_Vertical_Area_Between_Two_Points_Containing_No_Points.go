package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	points := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
	fmt.Println(maxWidthOfVerticalArea(points))
}

func maxWidthOfVerticalArea(points [][]int) int {
	res := make([]int, 0, len(points))
	for _, i := range points {
		res = append(res, i[0])
	}
	sort.Sort(sort.IntSlice(res))
	max := 0.0
	for i := 1; i < len(res); i++ {
		max = maxFunc(math.Abs(float64(res[i]-res[i-1])), max)
	}
	return int(max)
}

func maxFunc(val float64, max float64) float64 {
	if val > max {
		return val
	}
	return max
}
