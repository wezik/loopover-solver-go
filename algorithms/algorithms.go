package algorithms

import "loopover-solver/state"

type Algorithm interface {
	Apply(state *state.State) [][]int
}
