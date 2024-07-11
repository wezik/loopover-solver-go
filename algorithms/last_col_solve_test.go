package algorithms

import (
	"loopover-solver/state"
	"testing"
)

func TestLastColSolve(t *testing.T) {
	t.Run("Starting same and first of row cell", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 15},
			{9, 10, 11, 12},
			{13, 16, 8, 14},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}

		LastColSolve{}.Apply(state)
		got := state.Board[0:3]

		for i, row := range want {
			for j, cell := range row {
				if got[i][j] != cell {
					t.Errorf("got %v want %v", got, want)
				}
			}
		}
	})
	t.Run("All bottom row", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 13},
			{5, 6, 7, 15},
			{9, 10, 11, 14},
			{4, 16, 8, 12},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}

		LastColSolve{}.Apply(state)
		got := state.Board[0:3]

		for i, row := range want {
			for j, cell := range row {
				if got[i][j] != cell {
					t.Errorf("got %v want %v", got, want)
				}
			}
		}
	})
	t.Run("All last col", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 12},
			{9, 10, 11, 8},
			{13, 15, 14, 16},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}

		LastColSolve{}.Apply(state)
		got := state.Board[0:3]

		for i, row := range want {
			for j, cell := range row {
				if got[i][j] != cell {
					t.Errorf("got %v want %v", got, want)
				}
			}
		}
	})
}
