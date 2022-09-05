package main

import (
	"fmt"
)

func main() {
	s := "xyzzaz"
	// s := "aababcabc"
	fmt.Println(countGoodSubstrings(s))
}

func countGoodSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		str1 := ""
		mk := make(map[string]int)

		for j := i; j < len(s); j++ {
			str1 += string(s[j])
			mk[string(s[j])] = count
			if len(str1) == 3 && len(mk) == 3 {
				count += 1
			}
		}
	}
	return count
}
