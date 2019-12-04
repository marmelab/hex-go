package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"math"
)

const StartVertexId = 0

func buildGraphForPlayer1(stones []Stone, width int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	endVertexId := GetEndVertexId(width)

	graph.AddVertex(StartVertexId)
	addVertexToGraph(stones, graph, Player1)
	graph.AddVertex(endVertexId)

	addPlayer1StartArcsToGraph(stones, graph, width)
	addArcsToGraph(stones, graph, Player1)
	addPlayer1EndArcsToGraph(stones, graph, width, endVertexId)

	return graph
}

func buildGraphForPlayer2(stones []Stone, width int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	endVertexId := GetEndVertexId(width)

	graph.AddVertex(StartVertexId)
	addVertexToGraph(stones, graph, Player2)
	graph.AddVertex(endVertexId)

	addPlayer2StartArcsToGraph(stones, graph, width)
	addArcsToGraph(stones, graph, Player2)
	addPlayer2EndArcsToGraph(stones, graph, width, endVertexId)

	return graph
}

func GetEndVertexId(width int) int {
	return width*width + 1
}

func addPlayer1StartArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int) {
	for _, stone := range stones[:width] {
		// @todo: Check if there's another (clearer) way to do.
		distance := int64(math.Abs(float64(stone.player - 1)))

		if nil != graph.AddArc(StartVertexId, stone.id, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2StartArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int) {
	columnIndex := 1
	for i := range stones {
		if i == columnIndex {
			// @todo: Check if there's another (clearer) way to do.
			distance := int64(math.Abs(float64(stones[i-1].player - 2)))

			if nil != graph.AddArc(StartVertexId, stones[i-1].id, distance) {
				fmt.Errorf("error during Arc insertion")
			}

			columnIndex = columnIndex + width
		}
	}
}

func addPlayer1EndArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int, EndVertexId int) {
	offset := (width * width) - width
	for _, stone := range stones[offset:] {
		// @todo: Check if there's another (clearer) way to do.
		distance := int64(math.Abs(float64(stone.player - 1)))

		if nil != graph.AddArc(stone.id, EndVertexId, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2EndArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int, EndVertexId int) {
	columnIndex := width
	for i := range stones {
		if i == columnIndex {
			// @todo: Check if there's another (clearer) way to do.
			distance := int64(math.Abs(float64(stones[i-1].player - 2)))

			if nil != graph.AddArc(stones[i-1].id, EndVertexId, distance) {
				fmt.Errorf("error during Arc insertion")
			}
			columnIndex = columnIndex + width
		}
	}
}

func addArcsToGraph(stones []Stone, graph *dijkstra.Graph, player int) {
	for _, stone := range stones {
		if stone.player == player || stone.player == Empty {
			for _, neighbor := range getNeighborsForStone(stones, stone) {
				graph.AddArc(stone.id, neighbor.stone.id, int64(neighbor.distance))
			}
		}
	}
}

func addVertexToGraph(stones []Stone, graph *dijkstra.Graph, player int) {
	for _, stone := range stones {
		if stone.player == player || stone.player == Empty {
			graph.AddVertex(stone.id)
		}
	}
}
