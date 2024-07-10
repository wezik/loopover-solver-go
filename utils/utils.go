package utils

func Wrap(limit int, x int) int {
	result := x % limit
	if result < 0 {
		return limit + result
	} else {
		return result
	}
}

func CreateSolvedBoard(board [][]int) [][]int {
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

