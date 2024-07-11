package solver

import (
	"loopover-solver/algorithms"
	"loopover-solver/recorder"
	"loopover-solver/state"
)

func Solve(initialBoard [][]int) any {
	state := state.New(initialBoard)
	algorithm := algorithms.RowSolve{}
	algorithm.Apply(state)
	return recorder.FinishRecording()
}
