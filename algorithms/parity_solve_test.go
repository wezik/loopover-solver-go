package algorithms

import (
	"loopover-solver/state"
	"testing"
)

func TestParitySolve(t *testing.T) {
	t.Run("4x4 Default parity solve", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 16},
			{13, 14, 15, 12},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		ParitySolve{}.Apply(state)
		got := state.Board

		for i, row := range want {
			for j, cell := range row {
				if got[i][j] != cell {
					t.Errorf("got %v want %v", got, want)
					break
				}
			}
		}
	})
}
