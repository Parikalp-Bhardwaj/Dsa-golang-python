package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))
}

func stringMatching(words []string) []string {
	var res []string
	for _, i := range words {
		for _, j := range words {
			if i == j {
				continue
			}
			if strings.Contains(j, i) {
				res = append(res, i)
				break
			}
		}
	}
	return res
}
