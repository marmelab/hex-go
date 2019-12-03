package main

import (
	"testing"
)

func TestUserCanGetNeighborsForStone(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 2},
		{0, 1, 2, 2, 2},
		{0, 1, 0, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}

	stones := getStonesFromMatrix(matrix)

	stone := *NewStone(1, 1, PLAYER1)

	got := getNeighborsForStone(stones, stone)

	expected := []neighbor{
		*NewNeighbor(1, 0, 0),
		*NewNeighbor(2, 0, 0),
		*NewNeighbor(0, 1, 1),
		*NewNeighbor(0, 2, 1),
		*NewNeighbor(2, 2, 1),
	}

	for i, neighbor := range expected {
		if got[i] != neighbor {
			t.Errorf("got %v; want %v", got, neighbor)
		}
	}
}
