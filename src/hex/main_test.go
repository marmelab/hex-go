package main

import "github.com/RyanCarrier/dijkstra"

func bootstrapGraphPlayer1(matrix [][]int) (*dijkstra.Graph, int) {
	width := len(matrix)
	endVertexId := GetEndVertexId(width)

	stones := getStonesFromMatrix(matrix)

	return buildGraphForPlayer1(stones, width), endVertexId
}

func bootstrapGraphPlayer2(matrix [][]int) (*dijkstra.Graph, int) {
	width := len(matrix)
	endVertexId := GetEndVertexId(width)

	stones := getStonesFromMatrix(matrix)

	return buildGraphForPlayer2(stones, width), endVertexId
}
