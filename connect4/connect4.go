package main

import (
	"fmt"
	// "os"
	// "bufio"
)

type state struct {
	turn int
	board [][]int
	isOver bool
}

// intantiate new game state
func newGame() state {
	newState := state{} 
	newState.turn = 1
	newState.board = [][]int{
		[]int{0,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,0},
		[]int{0,0,0,0,0,0,0},
	}
	newState.isOver = false
	return newState
}

// display board tile
func (s state) displayTile(row int, column int) {
	switch s.board[row][column] {
	case 0:
		fmt.Printf(" - ")
		return
	case 1:
		fmt.Printf(" X ")
	case 2:
		fmt.Printf(" O ")
	}
	return
}

// print board to stdout
func (s state) displayBoard(showColumn bool) {
	if showColumn{
		fmt.Println(" 1  2  3  4  5  6  7 ")
	}
	
	for row := 0; row < 6; row++ {
		for column := 0; column < 7; column++ {
			s.displayTile(row, column)
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
    
	for (column > 7 || column < 1 || s.board[0][column - 1] != 0) {
		// flush input buffer
		var discard string
		fmt.Scanln(&discard)

		// get input again
		fmt.Printf("Invalid input. Player %d's move: ", s.turn)
		fmt.Scanln(&column)
	}
	// offset column by 1 for array
	return column - 1
}

// make player move
func (s *state) makeMove(column int) (int, int) {
	// look for first available tile in column
		for row := 5; row >= 0; row-- {
		if s.board[row][column] == 0 {
			// update tile
			s.board[row][column] = s.turn

			// update player turn
			s.turn = s.turn % 2 + 1

			// return tile position
			return row, column
		}
	}
	return -1, -1
}

// check for horizontal win
func (s state) checkHorizontal(row int) (bool) {
	for column, counter := 0, 0; column < 7; column++ {
		// increment or reset counter
		if s.board[row][column] == (s.turn % 2 + 1) {
			counter++
		} else {
			counter = 0
		}

		// check counter
		if counter == 4 {
			return true
		}
	}
	return false
}

// check for vertical win
func (s state) checkVertical(column int) (bool) {
	for row, counter := 0, 0; row < 6; row++ {
		// increment or reset counter
		if s.board[row][column] == (s.turn % 2 + 1) {
			counter++
		} else {
			counter = 0
		}

		// check counter
		if counter == 4 {
			return true
		}
	}
	return false
}

func main() {
	gameState := newGame()
	gameState.displayBoard(true)
	
	// REPL 
	for !gameState.isOver {
		move := gameState.getMove() 
		moveRow, moveColumn := gameState.makeMove(move)
		win := (gameState.checkHorizontal(moveRow) || gameState.checkVertical(moveColumn))
		
		gameState.displayBoard(true)
		println(win)
		println("-------------")
	}
}