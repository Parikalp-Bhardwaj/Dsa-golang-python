package main

import "fmt"

func main() {
	n := 5
	k := 2
	fmt.Println(findTheWinner(n, k))
}

func findTheWinner(n, k int) int {
	arr := make([]int, 0, n)
	for i := 1; i < n+1; i++ {
		arr = append(arr, i)
	}

	s := 0
	for len(arr) != 1 {
		rem := (s + k - 1) % len(arr)
		arr = append(arr[:rem], arr[rem+1:]...)
		s = rem
	}

	return arr[0] 
}
