package algorithms

import (
	"loopover-solver/state"
)

type LastColSolve struct{}

/*
[X][X][X][*]
[X][X][X][*]
[X][X][X][*]
[ ][ ][ ][ ]
* - Solved in this step
*/

func (s LastColSolve) Apply(state *state.State) {
	for destY := 0; destY < len(state.Board)-1; destY++ {
		lastX := len(state.Board[destY]) - 1
		lastY := len(state.Board) - 1
		currentY, currentX := state.FindCell(state.SolvedState[destY][lastX])
		if currentX == lastX && currentY != lastY {
			if destY == 0 {
				state.MoveVertical(lastX, lastY-1-currentY)
				continue
			}
			state.MoveVertical(lastX, lastY-currentY)
			cachedVertical := lastY - currentY
			state.MoveHorizontal(lastY, -1)
			state.MoveVertical(lastX, -cachedVertical)
			currentX = lastX - 1
		}
		state.MoveHorizontal(lastY, lastX-currentX)
		state.MoveVertical(lastX, -1)
	}
}
