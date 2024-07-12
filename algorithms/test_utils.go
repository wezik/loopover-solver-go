package algorithms

import "math/rand"

func compareRows(got, want []int) bool {
	for i, v := range want {
		if v != got[i] {
			return false
		}
	}
	return true
}

func shuffle(slice []int) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func createRandomBoard(rows int, cols int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}
	start := rand.Intn(100)
	nums := make([]int, rows*cols)
	for i := range nums {
		nums[i] = start + i
	}
	shuffle(nums)
	k := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = nums[k]
			k++
		}
	}
	return board
}

func createRandomBoardForCol(rows int, cols int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}
	start := rand.Intn(100)
	nums := make([]int, rows*cols)
	for i := range nums {
		nums[i] = start + i
	}
	edges := make([]int, rows+cols-1)
	k := 0
	e := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j == rows-1 || i == cols-1 {
				edges[e] = nums[k]
				e++
			}
			board[i][j] = nums[k]
			k++
		}
	}
	shuffle(edges)
	for i := 0; i < rows-1; i++ {
		board[i][cols-1] = edges[i]
	}
	for i := 0; i < cols; i++ {
		board[rows-1][i] = edges[rows-1+i]
	}
	return board
}
