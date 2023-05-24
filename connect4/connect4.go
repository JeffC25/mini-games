package main

import (
	"fmt"
	"math"
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
func (s state) displayTile(row int, col int) {
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
func (s state) displayBoard(showCol bool) {
	if showCol{
		fmt.Println(" 1  2  3  4  5  6  7 ")
	}
	
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			s.displayTile(row, col)
		}
		fmt.Println()
	}
}

// get move (row) from stdin
func (s state) getMove() int {
	// get input
	var col int
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
func (s *state) makeMove(col int) (int, int) {
	// look for first available tile in column
		for row := 5; row >= 0; row-- {
		if s.board[row][col] == 0 {
			// update tile
			s.board[row][col] = s.turn

			// update player turn
			s.turn = s.turn % 2 + 1

			// return tile position
			return row, col
		}
	}
	return -1, -1
}

// check for horizontal win
func (s state) checkHorizontal(row int) (bool) {
	for col, counter := 0, 0; col < 7; col++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn % 2 + 1) {
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
func (s state) checkVertical(col int) (bool) {
	for row, counter := 0, 0; row < 6; row++ {
		// increment or reset counter
		if s.board[row][col] == (s.turn % 2 + 1) {
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
func (s state) checkDiagonal (row int, col int) (bool) {
	startRowTopLeft := row - math.Min(row, col)
	startColTopLeft := col - math.Min(row, col)

	offset := 6 - math.Max(row, col)
	startRowBotRight := row + offset
	startColBotRight := col + offset 
}

func main() {
	gameState := newGame()
	gameState.displayBoard(true)
	
	// REPL 
	for !gameState.isOver {
		move := gameState.getMove() 
		moveRow, moveCol := gameState.makeMove(move)
		win := (gameState.checkHorizontal(moveRow) || gameState.checkVertical(moveCol))
		
		gameState.displayBoard(true)
		println(win)
		println("-------------")
	}
}