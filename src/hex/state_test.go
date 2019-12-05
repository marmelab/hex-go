package main

import (
	"testing"
)

func TestUserCanKnowIfHeCanWinInTheNextMove(t *testing.T) {

	matrix := [][]int{
		{0, 1, 0},
		{2, 0, 0},
		{2, 1, 0},
	}
	graph, endVertexId := bootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(StartVertexId, endVertexId)

	got := CanWinInOneTurn(bestPath)

	if got != true {
		t.Errorf("got %t; want true", got)
	}
}

func TestUserCanKnowIfHeCantWinInTheNextMove(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 0, 2},
		{2, 1, 0},
	}
	graph, endVertexId := bootstrapGraphPlayer2(matrix)

	bestPath, _ := graph.Shortest(StartVertexId, endVertexId)

	got := CanWinInOneTurn(bestPath)

	if got != true {
		t.Errorf("got %t; want true", got)
	}
}

func TestUserCanKnowIfTheGameIsWon(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 1, 0},
		{2, 1, 0},
	}
	graph, endVertexId := bootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(StartVertexId, endVertexId)

	got := isWon(bestPath)

	if got != true {
		t.Errorf("got %t; want true", got)
	}
}