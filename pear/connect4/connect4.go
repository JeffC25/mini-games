package main

import (
	"fmt"
	// "os"
	// "bufio"
)

type state struct {
	turn int
	board [][]string
}

func newGame() state {
	newState := state{} 
	newState.turn = 0
	newState.board = [][]string{
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
		[]string{" - ", " - ", " - ", " - ", " - ", " - ", " - "},
	}
	return newState
}

func (s state) displayBoard(showColumn bool) {
	if showColumn{
		fmt.Println(" 1  2  3  4  5  6  7 ")
	}
	
	for _, row := range s.board {
		for _, spot := range row {
			fmt.Print(spot)
		}
		fmt.Println()
	}
}

func (s state) getMove() {
	turn := s.turn % 2 + 1
	fmt.Printf("Player %d's turn: ", turn)

	var column int
	fmt.Scanln(&column)
    
	for (column > 7 || column < 1) {
		var discard string
		fmt.Scanln(&discard)

		fmt.Printf("Invalid input. Player %d's turn: ", turn)
		fmt.Scanln(&column)
	}
	fmt.Printf("Column: %d\n", column)
}

func main() {
	gameState := newGame()

	gameState.displayBoard(true)
	gameState.getMove() 
}