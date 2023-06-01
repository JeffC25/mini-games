package main

import "testing"

func TestHorizontal(t *testing.T) {
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
		for turn := 1; turn < 3; turn++ {
			testState.turn = int8(turn)
			if testState.checkHorizontal(int8(i)) {
				t.Error("False positive")
			}
		}
	}
}