package main

import "fmt"

func main() {
	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 1

	flowerbed := []int{0, 0, 1, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	i := 1
	mk := make([]int, len(flowerbed)+2)

	for j := 1; j < len(flowerbed)+1; j++ {
		mk[j] = flowerbed[j-1]
	}

	for i < len(mk)-1 {
		if mk[i-1] == 0 && mk[i] == 0 && mk[i+1] == 0 {
			mk[i] = 1
			n -= 1

		}
		if n <= 0 {
			return true
		}
		i += 1
	}

	return false

}
