package main

import "fmt"

func main() {
	plants := []int{2, 2, 3, 3}
	capacity := 5
	fmt.Println(wateringPlants(plants, capacity))
}

func wateringPlants(plants []int, capacity int) int {
	fill := capacity
	count := 0
	for i, p := range plants {
		if fill < p {
			fill = capacity
			count += 2 * i
		}
		fill -= p
		count += 1
	}
	return count
}
