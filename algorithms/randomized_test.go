//go:build random
// +build random

package algorithms

import (
	"loopover-solver/state"
	"math/rand"
	"testing"
)

func TestRandomized(t *testing.T) {
	t.Run("Randomized row solve test", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			size := rand.Intn(18) + 2
			board := createRandomBoard(size, size)
			state := state.New(board)

			want := state.SolvedState[0 : size-1]

			RowSolve{}.Apply(state)
			got := state.Board[0 : size-1]

			for i, row := range got {
				if !compareRows(row[0:size-1], want[i][0:size-1]) {
					t.Errorf("Got %v, wanted %v", row, want[i])
					return
				}
			}
		}
	})
	t.Run("Randomized last col solve test", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			size := rand.Intn(18) + 2
			board := createRandomBoardForCol(size, size)
			state := state.New(board)

			want := state.SolvedState[0 : size-1]

			LastColSolve{}.Apply(state)
			got := state.Board[0 : size-1]

			for i, row := range got {
				gotCell := row[len(row)-1]
				wantCell := want[i][len(want[i])-1]
				if gotCell != wantCell {
					t.Errorf("Got %v, wanted %v", gotCell, wantCell)
					return
				}
			}
		}
	})
}

func shuffle(slice []int) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func createRandomBoard(rows int, cols int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}
	start := rand.Intn(100)
	nums := make([]int, rows*cols)
	for i := range nums {
		nums[i] = start + i
	}
	shuffle(nums)
	k := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = nums[k]
			k++
		}
	}
	return board
}

func createRandomBoardForCol(rows int, cols int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}
	start := rand.Intn(100)
	nums := make([]int, rows*cols)
	for i := range nums {
		nums[i] = start + i
	}
	edges := make([]int, rows+cols-1)
	k := 0
	e := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == rows-1 || i == cols-1 {
				edges[e] = nums[k]
				e++
			}
			board[i][j] = nums[k]
			k++
		}
	}
	shuffle(edges)
	for i := 0; i < rows-1; i++ {
		board[i][cols-1] = edges[i]
	}
	for i := 0; i < cols; i++ {
		board[rows-1][i] = edges[rows-1+i]
	}
	return board
}
