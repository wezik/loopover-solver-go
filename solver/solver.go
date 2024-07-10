package solver

type State struct {
        board [][]int
        solvedState [][]int
        moves []string
}

func (r *State) recordMove(move string) {
        r.moves = append(r.moves, move)
}

func Solve(initialBoard [][]int) any {
        state := State{moves: []string{}}
        state.board = initialBoard
        state.solvedState = createSolvedState(initialBoard)
        return state.solvedState
}

func createSolvedState(board [][]int) [][]int {
        height := len(board)
        width := len(board[0])
        min := board[0][0]

        for _, row := range board {
                for _, cell := range row {
                        if cell < min {
                                min = cell
                        }
                }
        }

        result := make([][]int, height)
        for i := range result {
                result[i] = make([]int, width)
                for j := range result[i] {
                        result[i][j] = min
                        min++
                }
        }
        return result
}

