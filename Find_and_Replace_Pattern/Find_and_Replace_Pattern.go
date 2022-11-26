package main

import (
	"fmt"
)

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"

	fmt.Println(findAndReplacePattern(words, pattern))
}

func findAndReplacePattern(words []string, pattern string) []string {

	getPattern("deq", pattern)
	return words
}

func getPattern(w string, pattern string) {
	m1, m2 := make(map[string]string), make(map[string]string)

	for i := 0; i < len(pattern); i++ {
		b1 := w[i]
		b2 := w[i]
		if _, ok := m1[string(b1)]; !ok {
			m1[string(b1)] = string(b2)
		}
		if _, ok := m2[string(b2)]; !ok {
			m2[string(b2)] = string(b1)
		}

	}

	fmt.Println(m1, m2)

}
