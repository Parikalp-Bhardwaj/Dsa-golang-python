package main

import (
	"fmt"
)

func main() {
	// s := "aaabbb"
	s := "a"
	fmt.Println(checkString(s))
}

func checkString(s string) bool {
	f := -1
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "b" {
			f = i
			break
		}
	}

	if f == -1 {
		return true
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "a" {
			return i < f
		}
	}
	return true

}
