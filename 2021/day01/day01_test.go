package main

import (
	"testing"
)

func TestSolveDay01Part1(t *testing.T) {
	t.Run("Test SolveDay01Part1", func(t *testing.T) {

		tempList := []int{199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263}

		got := SolveDay01Part1(tempList)
		expected := 7

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
