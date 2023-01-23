package main

import "fmt"

func main(){
	m := 2
	n := 3
	indices := [][]int{{0,1},{1,1}}
	fmt.Println(oddCells(m,n,indices))
}

func oddCells(m int,n int,indices [][]int) int{
	mat := make([][]int,m)
	for i := 0; i < m;i++{
		mat[i]=make([]int,n)
	}

	for _,idx := range indices{
		row := idx[0]
		col := idx[1]
		for i:=0; i < n; i++{
			mat[row][i]++
		}
		for j:=0; j < m; j++{
			mat[j][col]++
		}
	}

	fmt.Println(mat)

	count := 0
	for _,j := range mat{
		for _,i := range j{
			if i % 2 != 0{
				count++
			}
		}
	}
	fmt.Println(count)
	return 0
}