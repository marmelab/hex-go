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

	got := getStonesFromMatrix(matrix)

	expected := *NewStone(1, 0, 0, 0)
	if got[0] != expected {
		t.Errorf("got %v; want %v", got, expected)
	}

	expected2 := *NewStone(25, 4, 4, Player2)
	if got[24] != expected2 {
		t.Errorf("got %v; want %v", got[24], expected2)
	}
}

func TestUserCanCalcDistanceBetweenEmptyCaseAndOwnedCase(t *testing.T) {
	stone2 := NewStone(1, 0, 0, Empty)

	got := getDistanceByTypeOfStone(Player1, *stone2)

	expected := 1

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}

func TestUserCanCalcDistanceBetweenTwoOwnedCase(t *testing.T) {
	stone2 := NewStone(1, 0, 0, Player1)

	got := getDistanceByTypeOfStone(Player1, *stone2)

	expected := 0

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}
}

func TestUserCanCalcDistanceBetweenOwnedCaseAndAnotherPlayerCase(t *testing.T) {
	stone2 := NewStone(1, 0, 0, Player2)

	got := getDistanceByTypeOfStone(Player1, *stone2)

	expected := -1

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
	stones := getStonesFromMatrix(matrix)

	got, err := loadStoneByCoord(stones, 2, 1)

	expected := *NewStone(8, 2, 1, 2)

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}

	if err != nil {
		t.Errorf("got %v; want nil", got)
	}
}

func TestUserCantGetAStoneWithWrongCoords(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 0, 2},
		{0, 1, 2, 2, 2},
		{0, 1, 0, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}
	stones := getStonesFromMatrix(matrix)

	got, err := loadStoneByCoord(stones, -1, -1)

	expected := Stone{}

	if got != expected {
		t.Errorf("got %v; want %v", got, expected)
	}

	if err == nil {
		t.Errorf("got %v; want not nil", err)
	}
}
