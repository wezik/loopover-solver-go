package algorithms

import "loopover-solver/state"

type LastRowSolve struct{}

/*
[X][X][X][X]
[X][X][X][X]
[X][X][X][X]
[*][*][*][**]
* - Solved in this step
** - Solved either in this step or possibly needs a parity fix if the board is even
*/

func (l LastRowSolve) Apply(state *state.State) {
	lastY := len(state.Board) - 1
	lastX := len(state.Board[lastY]) - 1

	_, currentX := state.FindCell(state.SolvedState[lastY][0])
	state.MoveHorizontal(lastY, lastX-1-currentX)
	for x := 1; x < lastX; x++ {
		currentY, currentX := state.FindCell(state.SolvedState[lastY][x])
		if currentY == lastY && currentX != lastX {
			state.MoveHorizontal(lastY, lastX-currentX)
			cacheHorizontal := lastX - currentX
			currentX = lastX
			state.MoveVertical(lastX, 1)
			state.MoveHorizontal(lastY, -cacheHorizontal)
			state.MoveVertical(lastX, -1)
		} else if currentY != lastY {
			state.MoveVertical(lastX, 1)
			state.MoveHorizontal(lastY, lastX-1-currentX)
			state.MoveVertical(lastX, -1)
			continue
		}
		state.MoveHorizontal(lastY, lastX-1-currentX)
	}

}
