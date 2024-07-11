package algorithms

import (
	"loopover-solver/state"
	"testing"
)

func compareRows(got, want []int) bool {
	for i, v := range want {
		if v != got[i] {
			return false
		}
	}
	return true
}

func TestRowSolve(t *testing.T) {
	t.Run("Starting same and first of row cell", func(t *testing.T) {
		board := [][]int{
			{6, 2, 11, 4},
			{5, 10, 7, 12},
			{1, 9, 3, 8},
		}
		state := state.New(board)
		wantRowAhead := []int{6, 2, 11}
		want := 5

		solveSingleCell(state, 1, 0)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][2]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}

	})
	t.Run("Starting same row and not first cell", func(t *testing.T) {
		board := [][]int{
			{1, 2, 11, 4},
			{7, 5, 6, 12},
			{9, 10, 3, 8},
		}
		state := state.New(board)
		wantRowAhead := []int{1, 2, 11}
		want := []int{5, 6, 7}

		solveSingleCell(state, 1, 2)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][0:3]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if !compareRows(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Starting in free fields cell", func(t *testing.T) {
		board := [][]int{
			{1, 2, 11, 4},
			{3, 5, 6, 12},
			{9, 10, 7, 8},
		}
		state := state.New(board)
		wantRowAhead := []int{1, 2, 11}
		want := []int{5, 6, 7}

		solveSingleCell(state, 1, 2)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][0:3]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if !compareRows(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Starting in last column cell", func(t *testing.T) {
		board := [][]int{
			{1, 2, 11, 4},
			{3, 5, 6, 12},
			{9, 10, 8, 7},
		}
		state := state.New(board)
		wantRowAhead := []int{1, 2, 11}
		want := []int{5, 6, 7}

		solveSingleCell(state, 1, 2)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][0:3]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if !compareRows(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Starting in the last column same row cell", func(t *testing.T) {
		board := [][]int{
			{1, 2, 11, 4},
			{3, 5, 6, 7},
			{9, 10, 8, 12},
		}
		state := state.New(board)
		wantRowAhead := []int{1, 2, 11}
		want := []int{5, 6, 7}

		solveSingleCell(state, 1, 2)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][0:3]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if !compareRows(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Solve row", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 10},
			{7, 5, 9, 8},
			{6, 4, 11, 12},
		}
		state := state.New(board)
		wantRowAhead := []int{1, 2, 3}
		want := []int{5, 6, 7}

		solveSingleRow(state, 1)
		gotRowAhead := state.Board[0][0:3]
		got := state.Board[1][0:3]

		if !compareRows(gotRowAhead, wantRowAhead) {
			t.Errorf("Got %v, wanted %v", gotRowAhead, wantRowAhead)
		}

		if !compareRows(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Solve row large", func(t *testing.T) {
		board := [][]int{
			{1, 2, 3, 4, 15},
			{6, 7, 8, 9, 5},
			{11, 12, 13, 14, 19},
			{16, 22, 25, 10, 20},
			{21, 17, 23, 24, 18},
		}
		state := state.New(board)
		want := [][]int{{1, 2, 3, 4}, {6, 7, 8, 9}, {11, 12, 13, 14}, {16, 17, 18, 19}}

		solveSingleRow(state, 3)
		got := state.Board[0:4]

		for i, row := range got[0:4] {
			if !compareRows(row, want[i]) {
				t.Errorf("Got %v, wanted %v", row, want[i])
			}
		}
	})
	t.Run("Solve row functional test", func(t *testing.T) {
		board := [][]int{
			{8, 23, 2, 9, 16},
			{13, 1, 17, 10, 4},
			{25, 12, 21, 15, 22},
			{18, 3, 24, 7, 5},
			{19, 6, 11, 20, 14},
		}
		state := state.New(board)
		want := [][]int{{1, 2, 3, 4}, {6, 7, 8, 9}, {11, 12, 13, 14}, {16, 17, 18, 19}}

		RowSolve{}.Apply(state)
		got := state.Board[0:4]

		for i, row := range got[0:4] {
			if !compareRows(row, want[i]) {
				t.Errorf("Got %v, wanted %v", row, want[i])
			}
		}
	})
}
