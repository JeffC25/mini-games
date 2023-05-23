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

// intantiate new game state
func newGame() state {
	newState := state{} 
	newState.turn = 1
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

// print board to stdout
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

// get move (row) from stdin
func (s state) getMove() int {
	// get input
	var column int
	fmt.Printf("Player %d's move: ", s.turn)
	fmt.Scanln(&column)
    
	for (column > 7 || column < 1 || s.board[0][column - 1] != " - ") {
		// flush input buffer
		var discard string
		fmt.Scanln(&discard)

		// get input again
		fmt.Printf("Invalid input. Player %d's move: ", s.turn)
		fmt.Scanln(&column)
	}
	// fmt.Printf("Column: %d\n", column)
	return column
}

func (s state) makeMove(column int) {


	// toggle player turn
	s.turn = (s.turn % 2) + 1
}

func main() {
	gameState := newGame()

	gameState.displayBoard(true)
	gameState.getMove() 
}