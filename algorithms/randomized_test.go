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
