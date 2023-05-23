package main

import (
	"fmt"
)

type state struct {
	turn int
	board [][]string
}

func newGame() state {
	newState := state{} 
	newState.turn = 0
	newState.board = [][]string{
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
		[]string{"- ", "- ", "- ", "- ", "- ", "- ", "- "},
	}
	return newState
}

func (s state) displayBoard() {
	fmt.Println("test")
}

func main() {
	gameState := newGame()

	gameState.displayBoard()
}