package main

import "fmt"

func main() {
	board := [][]string{{"X", ".", ".", "X"}, {".", ".", ".", "X"}, {".", ".", ".", "X"}}
	fmt.Println(countBattleships(board))

}

func countBattleships(board [][]string) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "." {
				continue
			}

			if i > 0 && board[i-1][j] == "X" {
				continue
			}
			if j > 0 && board[i][j-1] == "X" {
				continue
			}
			res++
		}
	}

	return res
}
