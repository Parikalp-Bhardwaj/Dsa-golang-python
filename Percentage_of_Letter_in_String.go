package main

import (
	"fmt"
)

func main() {
	s := "foobar"
	letter := "o"
	fmt.Println(percentageLetter(s, letter))
}

func percentageLetter(s string, letter string) int {
	var count int = 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case letter:
			count++

		}
	}

	return int(float64(count) / float64(len(s)) * 100)

}
