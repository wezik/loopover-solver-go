package solver

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
                testState.moveHorizontal(0, 6)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveHorizontalNegative", func(t *testing.T) {
                testState := State{Board: [][]int{{9, 13, 15, 18}, {12, 10, 11, 17}, {16, 14, 19, 20}}}
                want := [][]int{{9, 13, 15, 18}, {10, 11, 17, 12}, {16, 14, 19, 20}}
                testState.moveHorizontal(1, -5)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveHorizontalLengthOfRow", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
                testState.moveHorizontal(0, 4)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveHorizontalLengthOfRowNegative", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
                testState.moveHorizontal(1, -4)
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
                testState.moveVertical(3, 1)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveVerticalLenghtOfCol", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
                testState.moveVertical(3, 3)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveVerticalNegative", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := [][]int{{12, 10, 19, 17}, {9, 13, 11, 18}, {16, 14, 15, 20}}
                testState.moveVertical(2, -2)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
        t.Run("MoveVerticalNegativeLengthOfCol", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
                testState.moveVertical(2, -3)
                got := testState.Board
                if !compareBoards(got, want) {
                        t.Errorf("got %v want %v", got, want)
                }
        })
}

func TestRecordMovement(t *testing.T) {
        t.Run("RecordMovementRight", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := []string{"R0", "R0"}
                testState.moveVertical(0, 2)
                got := testState.Moves
                if len(got) != len(want) {
                        t.Errorf("got %v want %v", got, want)
                }
                for i, move := range want {
                        if move != got[i] {
                                t.Errorf("got %v want %v", got, want)
                        }
                }
        })
        t.Run("RecordMovementLeft", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := []string{"L1", "L1"}
                testState.moveVertical(1, -2)
                got := testState.Moves
                if len(got) != len(want) {
                        t.Errorf("got %v want %v", got, want)
                }
                for i, move := range want {
                        if move != got[i] {
                                t.Errorf("got %v want %v", got, want)
                        }
                }
        })
        t.Run("RecordMovementDown", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := []string{"D2", "D2"}
                testState.moveHorizontal(2, 2)
                got := testState.Moves
                if len(got) != len(want) {
                        t.Errorf("got %v want %v", got, want)
                }
                for i, move := range want {
                        if move != got[i] {
                                t.Errorf("got %v want %v", got, want)
                        }
                }
        })
        t.Run("RecordMovementUp", func(t *testing.T) {
                testState := State{Board: [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}}
                want := []string{"U0"}
                testState.moveHorizontal(0, -1)
                got := testState.Moves
                if len(got) != len(want) {
                        t.Errorf("got %v want %v", got, want)
                }
                for i, move := range want {
                        if move != got[i] {
                                t.Errorf("got %v want %v", got, want)
                        }
                }
        })
}

