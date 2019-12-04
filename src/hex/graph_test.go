package main

import (
	"github.com/RyanCarrier/dijkstra"
	"testing"
)

func TestUserCanBuildAGraphBasedOnStonesArray(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0},
		{2, 1, 0},
		{2, 1, 0},
	}

	stones := getStonesFromMatrix(matrix)
	graph := buildGraph(stones, Player1)

	path := []int{2, 5, 8, 9}
	expected := dijkstra.BestPath{Distance: int64(1), Path: path}

	got, _ := graph.Shortest(2, 9)

	if got.Distance != expected.Distance {
		t.Errorf("got %d; want %d", got.Distance, expected.Distance)
	}

}
