package main

import (
	"bufio"
	"fmt"
	"os"
)

func printBoard(board []string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println();
	}
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

	printBoard(board)
}