package main

import (
	"fmt"
	"strconv"
)

func main() {
	expression := "2-1-1"
	fmt.Println(diffWaysToCompute(expression))
}

func diffWaysToCompute(expression string) []int {
	output := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' {
			left := diffWaysToCompute(string(expression[:i]))
			right := diffWaysToCompute(string(expression[i+1:]))
			for _, x := range left {
				for _, y := range right {
					if expression[i] == '+' {
						output = append(output, x+y)
					} else if expression[i] == '-' {
						output = append(output, x-y)
					} else {
						output = append(output, x*y)
					}
				}
			}

		}
	}

	if len(output) == 0 {
		x, _ := strconv.Atoi(expression)
		output = append(output, x)
	}
	return output
}
