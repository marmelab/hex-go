package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
)

const StartVertexId = 0

const WeightOwned = 0
const WeightNotOwned = 1

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
		weight := getDistance(Player1, stone)
		if nil != graph.AddArc(StartVertexId, stone.id, weight) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2StartArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int) {
	columnIndex := 1
	for i := range stones {
		if i == columnIndex {
			currentStone := stones[i-1]
			weight := getDistance(Player2, currentStone)

			if nil != graph.AddArc(StartVertexId, currentStone.id, weight) {
				fmt.Errorf("error during Arc insertion")
			}

			columnIndex = columnIndex + width
		}
	}
}

func addPlayer1EndArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int, EndVertexId int) {
	offset := (width * width) - width
	for _, stone := range stones[offset:] {
		weight := getDistance(Player1, stone)

		if nil != graph.AddArc(stone.id, EndVertexId, weight) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2EndArcsToGraph(stones []Stone, graph *dijkstra.Graph, width int, EndVertexId int) {
	columnIndex := width
	for i := range stones {
		if i == columnIndex {
			currentStone := stones[i-1]
			weight := getDistance(Player2, currentStone)

			if nil != graph.AddArc(currentStone.id, EndVertexId, weight) {
				fmt.Errorf("error during Arc insertion")
			}
			columnIndex = columnIndex + width
		}
	}
}

func getDistance(player int, stone Stone) int64 {
	if player == stone.player {
		return WeightOwned
	}

	return WeightNotOwned
}

func addArcsToGraph(stones []Stone, graph *dijkstra.Graph, player int) {
	for _, stone := range stones {
		if stone.player == player || stone.player == Empty {
			for _, neighbor := range getNeighborsForStone(stones, stone, player) {
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
