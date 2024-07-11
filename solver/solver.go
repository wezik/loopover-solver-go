package solver

import (
	"loopover-solver/algorithms"
	"loopover-solver/recorder"
	"loopover-solver/state"
)

func Solve(initialBoard [][]int) any {
	state := state.New(initialBoard)
	algorithms.RowSolve{}.Apply(state)
	algorithms.LastColSolve{}.Apply(state)
	return recorder.FinishRecording()
}
