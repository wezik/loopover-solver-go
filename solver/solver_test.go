package solver

import (
	"loopover-solver/state"
	"testing"
)

func TestSolverFullSolve(t *testing.T) {
	t.Run("4x4 Default", func(t *testing.T) {
		board := [][]int{
			{10, 2, 15, 5},
			{6, 11, 16, 1},
			{12, 9, 3, 14},
			{7, 13, 4, 8},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		resolveState(state)
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
	t.Run("4x4 Default", func(t *testing.T) {
		board := [][]int{
			{7, 3, 11, 8},
			{16, 14, 1, 10},
			{4, 13, 2, 12},
			{9, 15, 6, 5},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		resolveState(state)
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
	t.Run("4x4 Default", func(t *testing.T) {
		board := [][]int{
			{8, 12, 3, 16},
			{4, 11, 6, 15},
			{1, 14, 10, 7},
			{13, 5, 2, 9},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		resolveState(state)
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
