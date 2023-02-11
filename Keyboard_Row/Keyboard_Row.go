package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

func findWords(words []string) []string {

	key := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	var res []string
	for i := 0; i < len(words); i++ {
		var str string
		l := words[i]
		for _, k := range key {
			for _, char := range words[i] {
				if strings.Contains(k, strings.ToLower(string(char))) {
					str += string(char)
				} else {
					break
				}
			}

		}

		if l == str {
			res = append(res, words[i])
		}

	}
	return res
}
