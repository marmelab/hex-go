package main

import (
	"testing"
)

func TestUserCanBuildAGridBasedOnMatrix(t *testing.T) {

	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 1, 0, 2},
		{0, 0, 1, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}

	got := buildFromMatrix(matrix)

	expected := *NewStone(0, 0, 0)
	if got[0] != expected {
		t.Errorf("got %v; want %v", got, expected)
	}

	expected2 := *NewStone(4, 4, PLAYER2)
	if got[24] != expected2 {
		t.Errorf("got %v; want %v", got[24], expected2)
	}
}

func TestUserCanCalcDistanceBetweenEmptyCaseAndOwnedCase(t *testing.T) {
	stone1 := NewStone(0, 0, PLAYER1)
	stone2 := NewStone(0, 0, EMPTY)

	got := getDistanceBetweenTwoStones(*stone1, *stone2)

	expected := 1

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}

func TestUserCanCalcDistanceBetweenTwoOwnedCase(t *testing.T) {
	stone1 := NewStone(0, 0, PLAYER1)
	stone2 := NewStone(0, 0, PLAYER1)

	got := getDistanceBetweenTwoStones(*stone1, *stone2)

	expected := 0

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}

func TestUserCanCalcDistanceBetweenOwnedCaseAndAnotherPlayerCase(t *testing.T) {
	stone1 := NewStone(0, 0, PLAYER1)
	stone2 := NewStone(0, 0, PLAYER2)

	got := getDistanceBetweenTwoStones(*stone1, *stone2)

	expected := -1

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}

func TestUserCanGetNeighborsForStone(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 2},
		{0, 1, 2, 2, 2},
		{0, 1, 0, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}

	stones := buildFromMatrix(matrix)

	stone := *NewStone(1, 1, PLAYER1)
	got := getNeighborsForStone(stones, stone)

	expected := [6]node{
		*NewNode(1, 0, 0),
		*NewNode(2, 0, 0),
		{},
		*NewNode(0, 1, 1),
		*NewNode(0, 2, 1),
		*NewNode(2, 2, 1),
	}

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}

}

func TestUserCanLoadStoneWithCoords(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 2},
		{0, 1, 2, 2, 2},
		{0, 1, 0, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}
	stones := buildFromMatrix(matrix)

	got := loadStoneByCoord(stones, 2, 1)

	expected := *NewStone(2, 1, 2)

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}
