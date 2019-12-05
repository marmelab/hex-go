package state

import (
	Graph "hex/graph"
	"testing"
)

func TestUserCanKnowIfHeCanWinInTheNextMove(t *testing.T) {

	matrix := [][]int{
		{0, 1, 0},
		{2, 0, 0},
		{2, 1, 0},
	}
	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

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
	graph, endVertexId := Graph.BootstrapGraphPlayer2(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

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
	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	got := IsWon(bestPath)

	if got != true {
		t.Errorf("got %t; want true", got)
	}
}

func TestUserCanKnowIfHeIsClosestThanOpponent(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 0, 0},
		{2, 1, 0},
	}
	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)
	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	graphOpponent, endVertexIdOpponennt := Graph.BootstrapGraphPlayer2(matrix)
	bestPathOpponent, _ := graphOpponent.Shortest(Graph.StartVertexId, endVertexIdOpponennt)

	got := IsClosestAsOpponent(bestPath, bestPathOpponent)

	if got != true {
		t.Errorf("got %t; want true", got)
	}
}

func TestUserCanKnowIfHeIsNotClosestThanOpponent(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 2, 0},
		{2, 1, 0},
	}
	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)
	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	graphOpponent, endVertexIdOpponennt := Graph.BootstrapGraphPlayer2(matrix)
	bestPathOpponent, _ := graphOpponent.Shortest(Graph.StartVertexId, endVertexIdOpponennt)

	got := IsClosestAsOpponent(bestPath, bestPathOpponent)

	if got != false {
		t.Errorf("got %t; want false", got)
	}
}

func TestUserCanHaveAnAdviceForHisNextMove(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 0, 2},
		{2, 1, 0},
	}
	graph, endVertexId := Graph.BootstrapGraphPlayer2(matrix)
	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	got := AdviceForNextMove(bestPath, graph)
	expected := 5

	if got != expected {
		t.Errorf("got %d; want %d", got, expected)
	}
}

func TestUserCanCheckAnAlreadyWinningGame(t *testing.T) {

	matrix := [][]int{
		{1, 1, 0},
		{2, 2, 2},
		{2, 1, 0},
	}
	graph, endVertexId := Graph.BootstrapGraphPlayer2(matrix)
	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	got := AdviceForNextMove(bestPath, graph)
	expected := 0

	if got != expected {
		t.Errorf("got %d; want %d", got, expected)
	}
}
