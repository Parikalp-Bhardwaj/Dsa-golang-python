package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Baseball Game.
// Memory Usage: 2.7 MB, less than 16.98% of Go online submissions for Baseball Game.




import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
}




func calPoints(ops []string) int {

	ls := []int{}
	sum := 0
	for _, i := range ops {
		if i != "C" && i != "D" && i != "+" {
			if s, err := strconv.Atoi(i); err == nil {
				ls = append(ls, s)
			}
		} else if i == "C" {
			ls = ls[:len(ls)-1]
		} else if i == "D" {
			ls = append(ls, 2*ls[len(ls)-1])
		} else if i == "+" {
			ls = append(ls, ls[len(ls)-1]+ls[len(ls)-2])
		}
	}
	for _, j := range ls {
		sum += j
	}
	return sum
}
