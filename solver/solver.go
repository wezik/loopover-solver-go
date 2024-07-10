package solver

import "loopover-solver/utils"

type State struct {
	Board       [][]int
	solvedState [][]int
	Moves       []string
}

func (r *State) moveHorizontal(row int, distance int) {
	rowLength := len(r.Board[row])
	newRow := make([]int, rowLength)
	for i, cell := range r.Board[row] {
		newRow[utils.Wrap(rowLength, i+distance)] = cell
	}
	r.Board[row] = newRow
}

func (r *State) moveVertical(column int, distance int) {
	columnLength := len(r.Board)
	newBoard := make([][]int, columnLength)

	for i, row := range r.Board {
		newBoard[i] = make([]int, len(row))
	}

	for i, row := range r.Board {
		for j, cell := range row {
			if j == column {
				newBoard[utils.Wrap(columnLength, i + distance)][j] = cell
			} else {
				newBoard[i][j] = cell
			}
		}
	}
	r.Board = newBoard
}

func Solve(initialBoard [][]int) any {
	state := State{
		Moves: []string{},
		Board: initialBoard,
		solvedState: utils.CreateSolvedBoard(initialBoard),
	}
	return state.solvedState
}



