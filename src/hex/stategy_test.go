package main

import (
	"github.com/RyanCarrier/dijkstra"
	"testing"
)

func bootstrapGraphPlayer1(matrix [][]int) (*dijkstra.Graph, int) {
	width := len(matrix)
	endVertexId := GetEndVertexId(width)

	stones := getStonesFromMatrix(matrix)

	return buildGraphForPlayer1(stones, width), endVertexId
}

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
		{2, 0, 0},
		{2, 0, 0},
	}
	graph, endVertexId := bootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(StartVertexId, endVertexId)

	got := CanWinInOneTurn(bestPath)

	if got != false {
		t.Errorf("got %t; want false", got)
	}
}
