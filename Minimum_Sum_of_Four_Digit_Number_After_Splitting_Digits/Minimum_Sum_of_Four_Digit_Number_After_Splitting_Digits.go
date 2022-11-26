package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 2932
	fmt.Println(minimumSum(num))

}

func minimumSum(num int) int {
	num1 := []int{}
	for i := 0; i < 4; i++ {
		num1 = append(num1, num%10)
		num /= 10

	}
	sort.Ints(num1)
	var n1 string
	var n2 string

	for i := 0; i < len(num1); i++ {
		if i%2 == 0 {
			n1 += strconv.Itoa(num1[i])
		} else {
			n2 += strconv.Itoa(num1[i])
		}
	}

	n3 := 0
	if i, err := strconv.Atoi(n1); err == nil {
		n3 = i

	}

	n4 := 0

	if i, err := strconv.Atoi(n2); err == nil {
		n4 = i

	}

	return n3 + n4
}
