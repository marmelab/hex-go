package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"math"
)

const StartVertexId = 0

func buildGraph(stones []Stone, player int, width int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	endVertexId := width*width + 1

	graph.AddVertex(StartVertexId)
	addVertexToGraph(stones, graph, player)
	graph.AddVertex(endVertexId)

	addStartArcsToGraph(stones, graph, player, width)
	addArcsToGraph(stones, graph, player)
	addEndArcsToGraph(stones, graph, player, width, endVertexId)

	return graph
}

func addStartArcsToGraph(stones []Stone, graph *dijkstra.Graph, player int, width int) {
	if player == Player1 {
		for _, stone := range stones[:width] {
			// @todo: Check if there's another (clearer) way to do.
			distance := int64(math.Abs(float64(stone.player - 1)))

			if nil != graph.AddArc(StartVertexId, stone.id, distance) {
				fmt.Errorf("error during Arc insertion")
			}
		}
	}
}

func addEndArcsToGraph(stones []Stone, graph *dijkstra.Graph, player int, width int, EndVertexId int) {
	if player == Player1 {

		offset := (width * width) - width
		for _, stone := range stones[offset:] {
			// @todo: Check if there's another (clearer) way to do.
			distance := int64(math.Abs(float64(stone.player - 1)))

			if nil != graph.AddArc(stone.id, EndVertexId, distance) {
				fmt.Errorf("error during Arc insertion")
			}
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
