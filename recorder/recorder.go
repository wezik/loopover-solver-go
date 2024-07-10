package recorder

import (
	"fmt"
	"math"
)

var moves []string = []string{}
var lastMove string = ""
var consecutiveMoves int = 0

func record(move string) {
        if lastMove != move {
                if (lastMove != "") {
                        moves = append(moves, fmt.Sprint(consecutiveMoves) + lastMove)
                }
                lastMove = move
                consecutiveMoves = 0
        } else {
                lastMove = move
        }
        consecutiveMoves++
}

func ResetRecorder() {
        moves = []string{}
        lastMove = ""
        consecutiveMoves = 0
}

func FinishRecording() []string {
        if (lastMove != "") {
                moves = append(moves, fmt.Sprint(consecutiveMoves) + lastMove)
        }
        result := moves
        ResetRecorder()
        return result
}

func RecordHorizontal(row, distance int) {
        for i := 0; float64(i) < math.Abs(float64(distance)); i++ {
                if distance > 0 {
                        record("R" + fmt.Sprint(row))
                } else {
                        record("L" + fmt.Sprint(row))
                }
        }
}

func RecordVertical(col, distance int) {
        for i := 0; float64(i) < math.Abs(float64(distance)); i++ {
                if distance > 0 {
                        record("D" + fmt.Sprint(col))
                } else {
                        record("U" + fmt.Sprint(col))
                }
        }
}

