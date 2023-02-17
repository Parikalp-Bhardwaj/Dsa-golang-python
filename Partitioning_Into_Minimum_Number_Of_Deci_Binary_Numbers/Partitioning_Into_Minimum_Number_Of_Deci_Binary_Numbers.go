package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "32"
	fmt.Println(minPartitions(n))
}

func minPartitions(n string) int {
	ans := 0
	for _, val := range n {
		if int(val) > ans {
			ans = int(val)
		}
	}
	max, _ := strconv.Atoi(string(ans))
	return max
}
