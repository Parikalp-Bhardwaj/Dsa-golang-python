package main

import (
	"fmt"
	"strconv"
)

func main() {
	// firstWord := "acb"
	// secondWord := "cba"
	// targetWord := "cdb"

	firstWord := "aaa"
	secondWord := "a"
	targetWord := "aab"

	fmt.Println(isSumEqual(firstWord, secondWord, targetWord))

}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	d := make(map[string]int)
	j := 0
	for i := 97; i < 123; i++ {
		d[string(rune(i))] = j
		j++
	}
	f1 := addingChar(d, firstWord)
	s1 := addingChar(d, secondWord)
	t1 := addingChar(d, targetWord)
	total := f1 + s1

	if total == t1 {
		return true

	}
	return false
}

func addingChar(d map[string]int, words string) int {
	sum := ""
	for _, v := range words {
		sum += strconv.Itoa(d[string(v)])
	}

	s1, err := strconv.Atoi(sum)
	if err != nil {
		panic(err)
	}

	return s1
}
