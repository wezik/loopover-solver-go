package recorder

import (
	"testing"
)

func TestRecord(t *testing.T) {
        t.Run("RecordHorizontal", func(t *testing.T) {
                want := []string{"1R0", "1L0"}

                RecordHorizontal(0, 1)
                RecordHorizontal(0, -1)

                got := FinishRecording()
                if len(got) != len(want) {
                        t.Errorf("FinishRecording() = %v; want %v", got, want)
                }
                for i := range want {
                        if got[i] != want[i] {
                                t.Errorf("FinishRecording() = %v; want %v", got, want)
                        }
                }
        })
        t.Run("RecordVertical", func(t *testing.T) {
                want := []string{"1D0", "1U1"}

                RecordVertical(0, 1)
                RecordVertical(1, -1)

                got := FinishRecording()
                if len(got) != len(want) {
                        t.Errorf("FinishRecording() = %v; want %v", got, want)
                }
                for i := range want {
                        if got[i] != want[i] {
                                t.Errorf("FinishRecording() = %v; want %v", got, want)
                        }
                }
        })
        t.Run("RecordHorizontalAndVertical", func(t *testing.T) {
                want := []string{"1R0", "1D0", "1L0", "1U1"}

                RecordHorizontal(0, 1)
                RecordVertical(0, 1)
                RecordHorizontal(0, -1)
                RecordVertical(1, -1)

                got := FinishRecording()
                if len(got) != len(want) {
                        t.Errorf("FinishRecording() = %v; want %v", got, want)
                }
                for i := range want {
                        if got[i] != want[i] {
                                t.Errorf("FinishRecording() = %v; want %v", got, want)
                        }
                }
        })
        t.Run("RecordConsecutiveMoves", func(t *testing.T) {
                want := []string{"4R0", "2D0", "1L2", "2L3", "2U1"}

                RecordHorizontal(0, 1)
                RecordHorizontal(0, 3)
                RecordVertical(0, 1)
                RecordVertical(0, 1)
                RecordHorizontal(2, -1)
                RecordHorizontal(3, -2)
                RecordVertical(1, -1)
                RecordVertical(1, -1)

                got := FinishRecording()
                if len(got) != len(want) {
                        t.Errorf("FinishRecording() = %v; want %v", got, want)
                }
                for i := range want {
                        if got[i] != want[i] {
                                t.Errorf("FinishRecording() = %v; want %v", got, want)
                        }
                }
        })
}

