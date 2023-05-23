package main

import "fmt"

func main() {
	board := []string{"-------", "-------", "-------", "-------", "-------", "-------",}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println();
	}
}