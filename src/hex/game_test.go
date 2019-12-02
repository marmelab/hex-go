package main

import (
	"testing"
)

func TestAnUserCanBuildAGameBasedOnMatrix(t *testing.T) {

	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 1, 0, 2},
		{0, 0, 1, 0, 2},
		{1, 1, 1, 0, 2},
		{1, 0, 0, 0, 2},
	}

	got := buildFromMatrix(matrix)
	expected := *NewStone(0,0,0)

	if got[0] != expected {
		t.Errorf("got %v; want %v", &got, &expected)
	}

	got2 := buildFromMatrix(matrix)
	expected2 := *NewStone(4,4,PLAYER2)
	if got[24] != expected {
		t.Errorf("got %v; want %v", &got2, &expected2)
	}
}
