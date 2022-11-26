package main

import (
	"fmt"
)

func main() {
	word := "abcdefd"
	ch := "d"
	fmt.Println(reversePrefix(word, ch))
}

func reversePrefix(word string, ch string) string {
	for i, s := range word {
		if string(s) == string(ch) {
			return reverse(word[:i+1]) + word[i+1:]
		}
	}

	return word
}

func reverse(w1 string) (result string) {
	for _, s1 := range w1 {
		result = string(s1) + result
	}
	return
}
