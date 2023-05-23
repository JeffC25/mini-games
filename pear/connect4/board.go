package main

import (
	"fmt"
)



func printBoard(board []string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println();
	}
}


func modifyArrah(board *[]string) {
	(*board)[0][0] = 'x'
	out := []rune(board)
	out[0] = 'x'
	return out
}

func main() {
	board := []string{
		"-------", 
		"-------", 
		"-------", 
		"-------", 
		"-------", 
		"-------",
	}

	modifyArrah(&board)
	printBoard(board)
}