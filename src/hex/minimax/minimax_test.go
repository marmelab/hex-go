package minimax

import (
	"hex/grid"
	"testing"
)

func TestUserCanGenerateAllGridPossibilitiesForAPlayer(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	stones := grid.GetStonesFromMatrix(matrix)
	got := getAllPossibleGrids(stones, grid.Player1)

	expected := grid.Player1

	if got[0][0].Player != expected {
		t.Errorf("got %d; want %d", got[0][0], expected)
	}

	if got[5][6].Player != expected {
		t.Errorf("got %d; want %d", got[5][6].Player, expected)
	}

}
