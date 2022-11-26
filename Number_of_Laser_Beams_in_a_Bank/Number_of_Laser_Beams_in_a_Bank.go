package main

import "strings"

func main() {
	// bank := []string{"011001", "000000", "010100", "001000"}
	bank := []string{"000", "111", "000"}
	println(numberOfBeams(bank))
}

func numberOfBeams(bank []string) int {
	prev := 0
	out := 0
	for _, b1 := range bank {
		count := strings.Count(b1, "1")
		if count > 0 {
			out += count * prev
			prev = count
		}
	}
	return out
}
