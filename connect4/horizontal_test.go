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

	testState.board = [][]int8{
		[]int8{1, 2, 1, 2, 1, 2, 1},
		[]int8{1, 2, 1, 2, 1, 2, 1},
		[]int8{1, 2, 1, 2, 1, 2, 1},
		[]int8{1, 2, 1, 2, 1, 2, 1},
		[]int8{1, 2, 1, 2, 1, 2, 1},
		[]int8{1, 2, 1, 2, 1, 2, 1},
	}
	
	for i := 0; i < 6; i++ {
		for turn := 1; turn < 3; turn++ {
			testState.turn = int8(turn)
			if testState.checkHorizontal(int8(i)) {
				t.Error("False positive")
			}
		}
	}

	testState.board = [][]int8{
		[]int8{1, 1, 1, 0, 1, 1, 1},
		[]int8{0, 2, 0, 2, 0, 2, 0},
		[]int8{2, 1, 2, 1, 0, 0, 0},
		[]int8{0, 0, 0, 1, 1, 1, 2},
		[]int8{0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 0, 2, 0, 2, 0, 0},
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