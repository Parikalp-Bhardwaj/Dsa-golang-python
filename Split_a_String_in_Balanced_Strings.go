package main

import (
	"fmt"
)

func main() {
	s := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	var count int
	var j int
	for _, i := range s {
		if i == 'L' {
			j += 1
		} else {
			j -= 1
		}
		if j == 0 {
			count += 1
		}

	}
	return count

	// for _, char := range s {
	// 	switch char {
	// 	case 'L':
	// 		j++
	// 	case 'R':
	// 		j--
	// 	}
	// 	if j == 0 {
	// 		count++
	// 	}
	// }
	// return count
}
