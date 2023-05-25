package main

import (
	"fmt"
	"math"
	// "os"
	// "bufio"
)

type state struct {
	turn int8
	board [][]int8
}

// intantiate new game state
func newGame() state {
	newState := state{} 
	newState.turn = 1
	newState.board = [][]int8{
		[]int8{0,0,0,0,0,0,0},
		[]int8{0,0,0,0,0,0,0},
		[]int8{0,0,0,0,0,0,0},
		[]int8{0,0,0,0,0,0,0},
		[]int8{0,0,0,0,0,0,0},
		[]int8{0,0,0,0,0,0,0},
	}
	return newState
}

// display board tile
func (s state) displayTile(row int8, col int8) {
	switch s.board[row][col] {
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
func (s state) displayBoard(showCol bool, showBar bool) {
	if showCol{
		fmt.Println(" 1  2  3  4  5  6  7 ")
	}
	
	var row int8
	var col int8
	for row = 0; row < 6; row++ {
		for col = 0; col < 7; col++ {
			s.displayTile(row, col)
		}
		fmt.Println()
	}

	if showBar {
		fmt.Println("--------------------")
	}
}

// get move (row) from stdin
func (s state) getMove() int8 {
	// get input
	var col int8
	fmt.Printf("Player %d's move: ", s.turn)
	fmt.Scanln(&col)
    
	for (col > 7 || col < 1 || s.board[0][col - 1] != 0) {
		// flush input buffer
		var discard string
		fmt.Scanln(&discard)

		// get input again
		fmt.Printf("Invalid input. Player %d's move: ", s.turn)
		fmt.Scanln(&col)
	}
	// offset col by 1 for array
	return col - 1
}

// make player move
func (s state) makeMove(col int8) (int8, int8) {
	// look for first available tile in column
	var row int8
	for row = 5; row >= 0; row-- {
		if s.board[row][col] == 0 {
			// update tile
			s.board[row][col] = s.turn

			// return tile position
			return row, col
		}
	}
	return -1, -1
}

// update player turn
func (s *state) updateTurn() {
	s.turn = s.turn % 2 + 1
}

// check for horizontal win
func (s state) checkHorizontal(row int8) (bool) {
	for col, counter := 0, 0; col < 7; col++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn) {
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
func (s state) checkVertical(col int8) (bool) {
	for row, counter := 0, 0; row < 6; row++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn) {
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

// check for diagonal win
func (s state) checkDiagonal (row int8, col int8) (bool) {
	rowLT := row - int8(math.Min(float64(row), float64(col)))
	colLT := col - int8(math.Min(float64(row), float64(col)))
	for counter := 0; rowLT < 6 && colLT < 7; rowLT, colLT = rowLT + 1, colLT + 1 {
		if s.board[rowLT][colLT] == (s.turn) {
			counter++
		} else {
			counter = 0
		}

		if counter == 4 {
			return true
		}
	}

	// offset := 6 - math.Max(row, col)
	// rowRB := row + offset
	// colRB := col + offset 
	return false
}

func main() {
	gameState := newGame()
	gameState.displayBoard(true, true)
	
	// REPL 
	for {
		move := gameState.getMove() 
		moveRow, moveCol := gameState.makeMove(move)
		gameState.displayBoard(true, true)

		if (
			gameState.checkHorizontal(moveRow) || 
			gameState.checkVertical(moveCol) ||
			gameState.checkDiagonal(moveRow, moveCol)) {
			fmt.Printf("Winner: Player %d\n", gameState.turn)

			return
		}
			
		gameState.updateTurn()
	}
}