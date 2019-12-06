package graph

import (
	"github.com/RyanCarrier/dijkstra"
	HexGrid "hex/grid"
)

func BootstrapGraphPlayer1(matrix [][]int) (*dijkstra.Graph, int) {
	grid := HexGrid.GetGridFromMatrix(matrix)
	endVertexId := GetEndVertexId(grid.Width)

	return BuildGraphForPlayer1(grid), endVertexId
}

func BootstrapGraphPlayer2(matrix [][]int) (*dijkstra.Graph, int) {
	grid := HexGrid.GetGridFromMatrix(matrix)
	endVertexId := GetEndVertexId(grid.Width)

	return BuildGraphForPlayer2(grid), endVertexId
}
