package state

import (
	"fmt"
	"loopover-solver/recorder"
	"loopover-solver/utils"
)

type State struct {
	Board       [][]int
	SolvedState [][]int
}

func New(initialBoard [][]int) *State {
	return &State{
		Board:       initialBoard,
		SolvedState: utils.CreateSolvedBoard(initialBoard),
	}
}

func (s *State) String() string {
	str := "\nBoard: \n"
	for _, row := range s.Board {
		for _, cell := range row {
			str += "[ " + fmt.Sprint(cell) + " ]"
		}
		str += "\n"
	}
	str += "\nSolved: \n"
	for _, row := range s.SolvedState {
		for _, cell := range row {
			str += "[ " + fmt.Sprint(cell) + " ]"
		}
		str += "\n"
	}
	return str
}

func (s *State) MoveHorizontal(row int, distance int) {
	rowLength := len(s.Board[row])
	newRow := make([]int, rowLength)
	for i, cell := range s.Board[row] {
		newRow[utils.Wrap(rowLength, i+distance)] = cell
	}
	recorder.RecordHorizontal(row, distance)
	s.Board[row] = newRow
}

func (s *State) MoveVertical(column int, distance int) {
	columnLength := len(s.Board)
	newBoard := make([][]int, columnLength)

	for i, row := range s.Board {
		newBoard[i] = make([]int, len(row))
	}

	for i, row := range s.Board {
		for j, cell := range row {
			if j == column {
				newBoard[utils.Wrap(columnLength, i+distance)][j] = cell
			} else {
				newBoard[i][j] = cell
			}
		}
	}
	recorder.RecordVertical(column, distance)
	s.Board = newBoard
}

func (s *State) FindCell(value int) (int, int) {
	for y, row := range s.Board {
		for x, cell := range row {
			if cell == value {
				return y, x
			}
		}
	}
	return -1, -1
}
