package solver

import (
	"loopover-solver/algorithms"
	"loopover-solver/recorder"
	"loopover-solver/state"
)

func Solve(initialBoard [][]int) any {
	state := state.New(initialBoard)
	resolveState(state)
	result := recorder.FinishRecording()
	recorder.ResetRecorder()
	return result
}

func resolveState(state *state.State) {
	algorithms.MainRowSolve{}.Apply(state)
	algorithms.LastColSolve{}.Apply(state)
	algorithms.LastRowSolve{}.Apply(state)
	lastY := len(state.Board) - 1
	lastX := len(state.Board[lastY]) - 1
	desiredCell := state.SolvedState[lastY][lastX]
	currentY, _ := state.FindCell(desiredCell)
	if currentY != lastY {
		algorithms.ParitySolve{}.Apply(state)
	}
}
