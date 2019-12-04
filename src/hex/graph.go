package main

import (
	"github.com/RyanCarrier/dijkstra"
)

func buildGraph(stones []Stone, player int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()

	addVertexToGraph(stones, graph, player)
	addArcsToGraph(stones, graph, player)

	return graph
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
