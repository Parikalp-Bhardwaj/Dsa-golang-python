package main

import (
	"fmt"
)

func main() {

	// list1 := []string{"Shogun", "Piatti", "Tapioca Express", "Burger King", "KFC"}
	// list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	list1 := []string{"happy", "sad", "good"}
	list2 := []string{"sad", "happy", "good"}

	// list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	// list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	fmt.Println(findRestaurant(list1, list2))
}
func findRestaurant(list1, list2 []string) []string {
	l1 := make(map[string]int)

	for i, s := range list1 {
		l1[s] = i
	}
	out := 3000
	var res []string
	for i, s := range list2 {
		if _, ok := l1[s]; ok {
			sum := i + l1[s]
			if sum < out {
				out = sum
				res = []string{s}
			} else if out == sum {
				res = append(res, s)
			}

		}
	}

	return res
}
