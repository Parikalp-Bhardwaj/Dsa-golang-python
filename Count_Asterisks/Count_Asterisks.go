package main

import (
	"fmt"
)

func main() {
	s := "l|*e*et|c**o|*de|"

	fmt.Println(countAsterisks(s))
}

func countAsterisks(s string) int {

	count := 0
	j := 0
	for _, char := range s {
		switch char {
		case '|':
			j++
		}

		if char == '*' && j%2 == 0 {
			count++
		}
	}
	return count
}
