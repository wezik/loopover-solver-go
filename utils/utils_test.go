package utils

import "testing"

func TestWrap(t *testing.T) {
	t.Run("WrapPositive", func(t *testing.T) {
		result := Wrap(5, 7)
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})
	t.Run("WrapNegative", func(t *testing.T) {
		result := Wrap(5, -1)
		if result != 4 {
			t.Errorf("Expected 4, got %d", result)
		}
	})
	t.Run("WrapZero", func(t *testing.T) {
		result := Wrap(5, 0)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})
	t.Run("WrapPositiveLimit", func(t *testing.T) {
		result := Wrap(5, 5)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})
	t.Run("WrapNegativeLimit", func(t *testing.T) {
		result := Wrap(5, -5)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})
}

func TestCreateSolvedBoard(t *testing.T) {
        testBoard := [][]int{{12, 10, 11, 17}, {9, 13, 15, 18}, {16, 14, 19, 20}}
        want := [][]int{{9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
        got := CreateSolvedBoard(testBoard)
        for i, row := range want {
                for j, cell := range row {
                        if got[i][j] != cell {
                                t.Errorf("got %v want %v", got, want)
                        }
                }
        }
}

