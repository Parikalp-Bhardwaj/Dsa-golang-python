package main

import "fmt"

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkIfPangram(sentence))
}

func checkIfPangram(sentence string) bool {
	mk := make(map[string]bool)
	for _, s := range sentence {
		mk[string(s)] = true
	}
	if len(mk) == 26{
		return true
	}
	return false
}
