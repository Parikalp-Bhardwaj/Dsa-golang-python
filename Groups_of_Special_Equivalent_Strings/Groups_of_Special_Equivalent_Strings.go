package main

import "fmt"

func main() {
	words := []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}
	fmt.Println(numSpecialEquivGroups(words))
}

func numSpecialEquivGroups(words []string) int {
	mapping := map[[52]int]interface{}{}
	for _, word := range words {
		var arr [52]int
		for idx, w := range word {
			if idx%2 == 0 {
				arr[w-'a']++
			} else {
				arr[w-'a'+26]++
			}
		}
		mapping[arr] = nil
	}

	return len(mapping)
}
