package main

import (
	"github.com/RyanCarrier/dijkstra"
	"testing"
)

func TestUserCanBuildGraphForPlayer1(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0},
		{2, 1, 0},
		{2, 1, 0},
	}

	graph, _ := bootstrapGraphPlayer1(matrix)

	path := []int{2, 5, 8, 9}
	expected := dijkstra.BestPath{Distance: int64(1), Path: path}

	got, _ := graph.Shortest(2, 9)

	if got.Distance != expected.Distance {
		t.Errorf("got %d; want %d", got.Distance, expected.Distance)
	}

}

func TestUserCanDetectTheShortestPathForPlayer1(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0},
		{2, 1, 0},
		{2, 1, 0},
	}

	graph, endVertexId := bootstrapGraphPlayer1(matrix)

	path := []int{StartVertexId, 2, 5, 8, endVertexId}
	expected := dijkstra.BestPath{Distance: int64(0), Path: path}

	got, _ := graph.Shortest(StartVertexId, endVertexId)

	if got.Distance != expected.Distance {
		t.Errorf("got %d; want %d", got.Distance, expected.Distance)
	}

	for i, path := range expected.Path {
		{
			if path != got.Path[i] {
				t.Errorf("got %d; want %d", path, got.Path[i])
			}
		}
	}
}

func TestUserCanDetectTheShortestPathForPlayer2(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2},
		{1, 2, 0},
		{2, 1, 0},
	}

	graph, endVertexId := bootstrapGraphPlayer2(matrix)
	
	path := []int{StartVertexId, 7, 5, 3, endVertexId}
	expected := dijkstra.BestPath{Distance: int64(0), Path: path}

	got, _ := graph.Shortest(StartVertexId, endVertexId)

	if got.Distance != expected.Distance {
		t.Errorf("got %d; want %d", got.Distance, expected.Distance)
	}

	for i, path := range expected.Path {
		{
			if path != got.Path[i] {
				t.Errorf("got %d; want %d", path, got.Path[i])
			}
		}
	}
}
