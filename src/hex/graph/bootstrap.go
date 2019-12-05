package graph

import (
	"github.com/RyanCarrier/dijkstra"
	"hex/grid"
)

func BootstrapGraphPlayer1(matrix [][]int) (*dijkstra.Graph, int) {
	width := len(matrix)
	endVertexId := GetEndVertexId(width)

	stones := grid.GetStonesFromMatrix(matrix)

	return BuildGraphForPlayer1(stones, width), endVertexId
}

func BootstrapGraphPlayer2(matrix [][]int) (*dijkstra.Graph, int) {
	width := len(matrix)
	endVertexId := GetEndVertexId(width)

	stones := grid.GetStonesFromMatrix(matrix)

	return BuildGraphForPlayer2(stones, width), endVertexId
}
