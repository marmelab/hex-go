package graph

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"hex/grid"
)

const StartVertexId = 0

func BuildGraphForPlayer1(stones []grid.Stone, width int) *dijkstra.Graph {
	graph := initGraph(width)

	addVertexToGraph(stones, graph, grid.Player1)

	addPlayer1StartArcsToGraph(stones, graph, width)
	addArcsToGraph(stones, graph, grid.Player1)
	addPlayer1EndArcsToGraph(stones, graph, width)

	return graph
}

func BuildGraphForPlayer2(stones []grid.Stone, width int) *dijkstra.Graph {
	graph := initGraph(width)

	addVertexToGraph(stones, graph, grid.Player2)

	addPlayer2StartArcsToGraph(stones, graph, width)
	addArcsToGraph(stones, graph, grid.Player2)
	addPlayer2EndArcsToGraph(stones, graph, width)

	return graph
}

func initGraph(width int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	endVertexId := GetEndVertexId(width)

	graph.AddVertex(StartVertexId)
	graph.AddVertex(endVertexId)

	return graph
}

func GetEndVertexId(width int) int {
	return width*width + 1
}

func addPlayer1StartArcsToGraph(stones []grid.Stone, graph *dijkstra.Graph, width int) {
	for _, stone := range stones[:width] {
		distance := getDistance(grid.Player1, stone)
		if nil != graph.AddArc(StartVertexId, stone.Id, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2StartArcsToGraph(stones []grid.Stone, graph *dijkstra.Graph, width int) {
	columnIndex := 1
	for i := range stones {
		if i == columnIndex {
			currentStone := stones[i-1]
			distance := getDistance(grid.Player2, currentStone)

			if nil != graph.AddArc(StartVertexId, currentStone.Id, distance) {
				fmt.Errorf("error during Arc insertion")
			}

			columnIndex = columnIndex + width
		}
	}
}

func addPlayer1EndArcsToGraph(stones []grid.Stone, graph *dijkstra.Graph, width int) {
	EndVertexId := GetEndVertexId(width)
	offset := (width * width) - width
	for _, stone := range stones[offset:] {
		distance := getDistance(grid.Player1, stone)

		if nil != graph.AddArc(stone.Id, EndVertexId, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2EndArcsToGraph(stones []grid.Stone, graph *dijkstra.Graph, width int) {
	EndVertexId := GetEndVertexId(width)
	columnIndex := width
	for i := range stones {
		if i == columnIndex {
			currentStone := stones[i-1]
			distance := getDistance(grid.Player2, currentStone)

			if nil != graph.AddArc(currentStone.Id, EndVertexId, distance) {
				fmt.Errorf("error during Arc insertion")
			}
			columnIndex = columnIndex + width
		}
	}
}

func getDistance(player int, stone grid.Stone) int64 {
	if player == stone.Player {
		return grid.DistanceOwned
	}

	return grid.DistanceEmpty
}

func addArcsToGraph(stones []grid.Stone, graph *dijkstra.Graph, player int) {
	for _, stone := range stones {
		if stone.Player == player || stone.Player == grid.Empty {
			for _, neighbor := range grid.GetNeighborsForStone(stones, stone, player) {
				graph.AddArc(stone.Id, neighbor.Stone.Id, int64(neighbor.Distance))
			}
		}
	}
}

func addVertexToGraph(stones []grid.Stone, graph *dijkstra.Graph, player int) {
	for _, stone := range stones {
		if stone.Player == player || stone.Player == grid.Empty {
			graph.AddVertex(stone.Id)
		}
	}
}
