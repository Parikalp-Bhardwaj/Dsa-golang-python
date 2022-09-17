package main

import "fmt"

func main() {
	// s := "())"
	// s := "((())"
	s := "((("
	n := []int{3, 5, 6, 7, 8, 9}
	n = n[:len(n)-1]
	fmt.Println(n)
	fmt.Println(minAddToMakeValid(s))
}
func minAddToMakeValid(s string) int {
	// var res []string
	var p1 int
	var p2 int
	for _, s1 := range s {
		if s1 == '(' {
			p1++
		} else {
			if p1 == 0 {
				p2++
			} else {
				p1--
			}
		}
	}
	p2 += p1
	return p2
}
