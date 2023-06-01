package main

import (
	// "fmt"
	"testing"
)

func TestDiagonal(t *testing.T) {
	// set up
	testState := state{}
	testState.turn = 1

	// test false positive
	testState.board = [][]int8{
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
	}
	
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if testState.checkDiagonalLB(int8(i), int8(j)) || testState.checkDiagonalLT(int8(i), int8(j)) {
				t.Error("False positive")
			}
		}
	}

	// test diagonalLB
	testState.board = [][]int8{
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 1, 0, 0, 0},
		[]int8{0, 0, 1, 0, 0, 0, 0},
		[]int8{0, 1, 0, 0, 0, 0, 0},
		[]int8{1, 0, 0, 0, 0, 0, 0},
	}

	if !testState.checkDiagonalLB(5, 0) {
		t.Error("Should be diagonal win")
	}

	// test diagonalLT
	testState.board = [][]int8{
		[]int8{2, 0, 0, 0, 0, 0, 0},
		[]int8{0, 2, 0, 0, 0, 0, 0},
		[]int8{0, 0, 2, 0, 0, 0, 0},
		[]int8{0, 0, 0, 2, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0},
	}

	testState.updateTurn()
	if !testState.checkDiagonalLT(0, 0) {
		t.Error("Should be diagonal win")
	}

}
