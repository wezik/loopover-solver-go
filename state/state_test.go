package state

import "testing"

func compareBoards(got [][]int, want [][]int) bool {
	for i, row := range want {
		for j, cell := range row {
			if got[i][j] != cell {
				return false
			}
		}
	}
	return true
}

func TestMoveHorizontal(t *testing.T) {
	t.Run("MoveHorizontal", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{11, 17, 12, 10}, {9, 13, 15, 18}, {16, 14, 19, 20}}
		testState.MoveHorizontal(0, 6)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveHorizontalNegative", func(t *testing.T) {
		testState := State{Board: [][]int{{9, 13, 15, 18}, {12, 10, 11, 17}, {16, 14, 19, 20}}}
		want := [][]int{{9, 13, 15, 18}, {10, 11, 17, 12}, {16, 14, 19, 20}}
		testState.MoveHorizontal(1, -5)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveHorizontalLengthOfRow", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
		testState.MoveHorizontal(0, 4)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveHorizontalLengthOfRowNegative", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
		testState.MoveHorizontal(1, -4)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestMoveVertical(t *testing.T) {
	t.Run("MoveVertical", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 11, 20}, {9, 13, 15, 17}, {16, 14, 19, 18}}
		testState.MoveVertical(3, 1)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveVerticalLenghtOfCol", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
		testState.MoveVertical(3, 3)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveVerticalNegative", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 19, 17}, {9, 13, 11, 18}, {16, 14, 15, 20}}
		testState.MoveVertical(2, -2)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("MoveVerticalNegativeLengthOfCol", func(t *testing.T) {
		testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
		want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
		testState.MoveVertical(2, -3)
		got := testState.Board
		if !compareBoards(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestFindCell(t *testing.T) {
	t.Run("FindCell", func(t *testing.T) {
		testState := State{Board: [][]int{{5, 99, 12, 9}, {20, 25, 60, 9}}}
		wantRow := 0
		wantCol := 3
		gotRow, gotCol := testState.FindCell(9)
		if gotRow != wantRow || gotCol != wantCol {
			t.Errorf("got %v %v want %v %v", gotRow, gotCol, wantRow, wantCol)
		}
	})
	t.Run("FindCell", func(t *testing.T) {
		testState := State{Board: [][]int{{5, 99, 12, 9}, {20, 25, 60, 9}}}
		wantRow := 1
		wantCol := 1
		gotRow, gotCol := testState.FindCell(25)
		if gotRow != wantRow || gotCol != wantCol {
			t.Errorf("got %v %v want %v %v", gotRow, gotCol, wantRow, wantCol)
		}
	})
	t.Run("FindNonExistantCell", func(t *testing.T) {
		testState := State{Board: [][]int{{5, 99, 12, 9}, {20, 25, 60, 9}}}
		wantRow := -1
		wantCol := -1
		gotRow, gotCol := testState.FindCell(7)
		if gotRow != wantRow || gotCol != wantCol {
			t.Errorf("got %v %v want %v %v", gotRow, gotCol, wantRow, wantCol)
		}
	})
}
