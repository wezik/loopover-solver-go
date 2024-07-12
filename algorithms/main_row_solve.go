package algorithms

import (
	"loopover-solver/state"
)

type MainRowSolve struct{}

/*
[*][*][*][ ]
[*][*][*][ ]
[*][*][*][ ]
[ ][ ][ ][ ]
* - Solved in this step
*/

func (s MainRowSolve) Apply(state *state.State) {
	for y := 0; y < len(state.Board)-1; y++ {
		solveSingleRow(state, y)
	}
}

func solveSingleRow(state *state.State, destY int) {
	for x := 0; x < len(state.Board[destY])-1; x++ {
		solveSingleCell(state, destY, x)
	}
}

func solveSingleCell(state *state.State, destY, destX int) {
	lastX := len(state.Board[0]) - 1
	currentY, currentX := state.FindCell(state.SolvedState[destY][destX])
	if currentY == destY && currentX != lastX {
		if destX == 0 {
			state.MoveHorizontal(currentY, lastX-1-currentX)
			return
		}
		state.MoveVertical(currentX, 1)
		currentY = currentY + 1
		state.MoveHorizontal(currentY, lastX-currentX)
		state.MoveVertical(currentX, -1)
		currentX = lastX
		state.MoveVertical(currentX, destY-currentY)
	} else if currentY != destY {
		state.MoveHorizontal(currentY, lastX-currentX)
		currentX = lastX
		state.MoveVertical(currentX, destY-currentY)
	}
	state.MoveHorizontal(destY, (lastX-1)-currentX)
}
