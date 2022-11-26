package main

import "fmt"

func main() {
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))

}

func removeOuterParentheses(s string) string {
	var res string
	c := 0
	for _, v := range s {
		if string(v) == "(" {
			c++
			if c > 1 {
				res += "("
			}
		} else {
			if c > 1 {
				res += ")"
			}
			c -= 1
		}
	}
	return res
}
