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

	stone := *NewStone(7, 1, 1, Player1)

	got := getNeighborsForStone(stones, stone, Player1)

	expected := []Neighbor{
		*NewNeighbor(*NewStone(2, 1, 0, Player1), 0),
		*NewNeighbor(*NewStone(3, 2, 0, Player1), 0),
		*NewNeighbor(*NewStone(6, 0, 1, Empty), 1),
		*NewNeighbor(*NewStone(11, 0, 2, Empty), 1),
		*NewNeighbor(*NewStone(12, 1, 2, Player1), 0),
	}

	for i, neighbor := range expected {
		if got[i] != neighbor {
			t.Errorf("got %v; want %v", got[i], neighbor)
		}
	}
}

func TestUserCanGetNeighborsForStoneInEdge(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 2},
		{0, 1, 2, 2, 2},
		{0, 1, 0, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}

	stones := getStonesFromMatrix(matrix)

	stone := *NewStone(1, 0, 0, Player1)

	got := getNeighborsForStone(stones, stone, Player1)

	expected := []Neighbor{
		*NewNeighbor(*NewStone(2, 1, 0, Player1), 0),
		*NewNeighbor(*NewStone(6, 0, 1, Empty), 1),
	}

	for i, neighbor := range expected {
		if got[i] != neighbor {
			t.Errorf("got %v; want %v", got[i], neighbor)
		}
	}
}

func TestUserCanGetNeighborsForAPlayer2Stone(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0},
		{0, 2, 2, 2},
		{2, 1, 0, 0},
		{1, 1, 1, 0},
	}

	stones := getStonesFromMatrix(matrix)

	stone := *NewStone(6, 1, 1, Player2)

	got := getNeighborsForStone(stones, stone, Player2)

	expected := []Neighbor{
		*NewNeighbor(*NewStone(7, 2, 1, Player2), 0),
		*NewNeighbor(*NewStone(5, 0, 1, Empty), 1),
		*NewNeighbor(*NewStone(9, 0, 2, Player2), 0),
	}

	for i, neighbor := range expected {
		if got[i] != neighbor {
			t.Errorf("got %v; want %v", got[i], neighbor)
		}
	}
}
