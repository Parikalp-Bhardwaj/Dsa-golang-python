package main

import (
	"fmt"
	"strings"
)

func main() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBa"
	fmt.Println(camelMatch(queries, pattern))
}

func camelMatch(queries []string, pattern string) []bool {
	mk := make([]bool, 0)
	for _, query := range queries {
		count := 0
		fl := 0
		for i := 0; i < len(query); i++ {
			if count < len(pattern) {
				if string(pattern[count]) == string(query[i]) {
					count += 1
				} else if string(query[i]) == strings.ToUpper(string(query[i])) {
					fl += 1
					break
				}
			} else if strings.Compare(string(query[i:]), strings.ToLower(string(query[i:]))) != 0 {
				fl += 1
				break
			} else {
				break
			}
		}

		if fl == 0 && count == len(pattern) {
			mk = append(mk, true)
		} else {
			mk = append(mk, false)
		}

	}

	return mk
}
