package minimax

import (
	HexGrid "hex/grid"
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

	grid := HexGrid.GetGridFromMatrix(matrix)
	got := getAllPossibleGrids(grid, HexGrid.Player1)

	expected := HexGrid.Player1

	firstGrid := got[0]
	if firstGrid.Stones[0].Player != expected {
		t.Errorf("got %d; want %d", firstGrid.Stones[0].Player, expected)
	}

	sixth := got[5]
	if sixth.Stones[6].Player != expected {
		t.Errorf("got %d; want %d", sixth.Stones[6].Player, expected)
	}
}
