package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	strs := [][]string{{"alic3", "bob", "3", "4", "00000"},
		{"5232", "yv", "22", "c", "yawgs", "928", "4003", "2"},
		{"5", "5", "iyid", "56204", "41", "b6d", "383", "ic"},
		{"1", "01", "001", "0001"}}
	for _, i := range strs {
		fmt.Println(maximumValue(i))

	}
	
}

func maximumValue(strs []string) int {
	max := 0
	for _, str := range strs {
		isDigitOnly := true
		for _, char := range str {
			if !unicode.IsDigit(char) {
				isDigitOnly = false
				break
			}
		}

		if isDigitOnly {
			val, _ := strconv.Atoi(str)
			if val > max {
				max = val
			}
		} else {
			length := len(str)
			if length > max {
				max = length
			}
		}
	}
	return max

}

func maxVal(count int, j string) int {
	val, _ := strconv.Atoi(j)
	if count > val {
		return count
	}
	return val
}
