package main

import "fmt"

func main() {
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	fmt.Println(findJudge(n, trust))
}

func findJudge(n int, trust [][]int) int {
	ans := make([]int,n+1)
	for _,x := range trust{
		ans[x[0]] -=1
		ans[x[1]] +=1
	}

	for i := 1; i <= n;i++{
		if (ans[i] == n-1){
			return i
		}
	}
	return -1
}
