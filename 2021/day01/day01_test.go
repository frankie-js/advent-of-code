package main

import (
	"testing"
)

func TestSolveDay01Part1(t *testing.T) {
	t.Run("Test SolveDay01Part1", func(t *testing.T) {

		got := SolveDay01Part1([]int{199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263})
		expected := 7

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

// func TestSolveDay01Part2(t *testing.T) {
// 	t.Run("Test SolveDay01Part2", func(t *testing.T) {
// 		got := SolveDay01Part2([]int{199,
// 			200,
// 			208,
// 			210,
// 			200,
// 			207,
// 			240,
// 			269,
// 			260,
// 			263})
// 		expected := 5
// 		if got != expected {
// 			t.Errorf("expected '%d' but got '%d'", expected, got)
// 		}
// 	})
// }

func TestSolveDay01Part2(t *testing.T) {
	t.Run("Test SolveDay01Part2", func(t *testing.T) {
		got := SolveDay01Part2(`199
200
208
210
200
207
240
269
260
263`)
		expected := 5
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
