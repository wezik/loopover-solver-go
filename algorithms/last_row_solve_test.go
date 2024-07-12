package algorithms

import (
	"loopover-solver/state"
	"testing"
)

func TestLastRowSolve(t *testing.T) {
	t.Run("Standard 4x4", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 16, 15, 14},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, -1},
			{13, 14, 15, -1},
		}

		LastRowSolve{}.Apply(state)
		got := state.Board

		for i, row := range want {
			for j, cell := range row {
				if i >= len(want)-2 && j == len(row)-1 {
					continue
				}
				if got[i][j] != cell {
					t.Errorf("got %v want %v", got, want)
					break
				}
			}
		}
	})
	t.Run("Standard 2x2", func(t *testing.T) {
		board := [][]int{
			{1, 2},
			{4, 3},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2},
			{3, 4},
		}

		LastRowSolve{}.Apply(state)
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
	t.Run("Standard 3x3", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{8, 9, 7},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		LastRowSolve{}.Apply(state)
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
	t.Run("Cell in last column and not last row", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 24},
			{20, 21, 22, 23, 25},
		}
		state := state.New(board)
		want := [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}

		LastRowSolve{}.Apply(state)
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
