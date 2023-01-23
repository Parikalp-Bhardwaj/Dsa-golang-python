package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))

	names = []string{"Alice", "Bob", "Bob"}
	heights = []int{155, 185, 150}
	fmt.Println(sortPeople(names, heights))
	names = []string{"Alice", "Bob", "Bob"}
	heights = []int{155, 185, 150}

	fmt.Println(sortPeople2(names, heights))

}

func sortPeople(names []string, heights []int) []string {
	res := make(map[int]string)
	for i := 0; i < len(names); i++ {
		res[heights[i]] = names[i]

	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[j] < heights[i]
	})
	fmt.Println(heights)

	ans := make([]string, 0, len(res))

	for j := 0; j < len(res); j++ {
		ans = append(ans, res[heights[j]])
	}

	return ans
}

func sortPeople2(names []string, heights []int) int {
	type people struct {
		name   string
		height int
	}
	res := make([]people, 0)
	for i := 0; i < len(names); i++ {
		ppl := people{name: names[i], height: heights[i]}
		res = append(res, ppl)
	}
	fmt.Println(res)
	sort.Slice(res, func(i, j int) bool {
		return res[i].height > res[j].height
	})
	fmt.Println(res)

	ans := make([]string, 0, len(names))
	for i := 0; i < len(names); i++ {
		ans = append(ans, res[i].name)
	}
	fmt.Println(ans)
	return 0
}
