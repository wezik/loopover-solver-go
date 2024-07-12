package algorithms

import "loopover-solver/state"

type ParitySolve struct{}

func (s ParitySolve) Apply(state *state.State) {
	lastY := len(state.Board) - 1
	lastX := len(state.Board[lastY]) - 1
	desiredCell := state.SolvedState[lastY][lastX]
	currentY, _ := state.FindCell(desiredCell)
	if currentY != lastY {
		for i := 0; i < lastY; i++ {
			state.MoveVertical(lastX, 1)
			state.MoveHorizontal(lastY, -1)
			state.MoveVertical(lastX, -1)
			if i < lastY-1 {
				state.MoveHorizontal(lastY, -1)
			}
		}
	}
}
