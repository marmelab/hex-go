package graph

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	HexGrid "hex/grid"
)

const StartVertexId = 0

func BuildGraphForPlayer(grid HexGrid.Grid, player int) *dijkstra.Graph {

	var graph *dijkstra.Graph
	if player == HexGrid.Player1 {
		graph = BuildGraphForPlayer1(grid)
	} else if player == HexGrid.Player2 {
		graph = BuildGraphForPlayer2(grid)
	}

	return graph
}

func BuildGraphForPlayer1(grid HexGrid.Grid) *dijkstra.Graph {
	graph := initGraph(grid.Width)

	addVertexToGraph(grid, graph, HexGrid.Player1)

	addPlayer1StartArcsToGraph(grid, graph)
	addArcsToGraph(grid, graph, HexGrid.Player1)
	addPlayer1EndArcsToGraph(grid, graph)

	return graph
}

func BuildGraphForPlayer2(grid HexGrid.Grid) *dijkstra.Graph {
	graph := initGraph(grid.Width)

	addVertexToGraph(grid, graph, HexGrid.Player2)

	addPlayer2StartArcsToGraph(grid, graph)
	addArcsToGraph(grid, graph, HexGrid.Player2)
	addPlayer2EndArcsToGraph(grid, graph)

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

func addPlayer1StartArcsToGraph(grid HexGrid.Grid, graph *dijkstra.Graph) {
	for _, stone := range grid.Stones[:grid.Width] {
		distance := getDistance(HexGrid.Player1, stone)
		if nil != graph.AddArc(StartVertexId, stone.Id, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2StartArcsToGraph(grid HexGrid.Grid, graph *dijkstra.Graph) {
	columnIndex := 1
	for i := range grid.Stones {
		if i == columnIndex {
			currentStone := grid.Stones[i-1]
			distance := getDistance(HexGrid.Player2, currentStone)

			if nil != graph.AddArc(StartVertexId, currentStone.Id, distance) {
				fmt.Errorf("error during Arc insertion")
			}

			columnIndex = columnIndex + grid.Width
		}
	}
}

func addPlayer1EndArcsToGraph(grid HexGrid.Grid, graph *dijkstra.Graph) {
	EndVertexId := GetEndVertexId(grid.Width)
	offset := (grid.Width * grid.Width) - grid.Width
	for _, stone := range grid.Stones[offset:] {
		distance := getDistance(HexGrid.Player1, stone)

		if nil != graph.AddArc(stone.Id, EndVertexId, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2EndArcsToGraph(grid HexGrid.Grid, graph *dijkstra.Graph) {
	EndVertexId := GetEndVertexId(grid.Width)
	columnIndex := grid.Width
	for i := range grid.Stones {

		if i+1 == columnIndex {
			currentStone := grid.Stones[i]
			distance := getDistance(HexGrid.Player2, currentStone)

			if nil != graph.AddArc(currentStone.Id, EndVertexId, distance) {
				fmt.Errorf("error during Arc insertion")
			}
			columnIndex = columnIndex + grid.Width
		}
	}
}

func getDistance(player int, stone HexGrid.Stone) int64 {
	if player == stone.Player {
		return HexGrid.DistanceOwned
	}

	return HexGrid.DistanceEmpty
}

func addArcsToGraph(grid HexGrid.Grid, graph *dijkstra.Graph, player int) {
	for _, stone := range grid.Stones {
		if stone.Player == player || stone.Player == HexGrid.Empty {
			for _, neighbor := range HexGrid.GetNeighborsForStone(grid, stone, player) {
				graph.AddArc(stone.Id, neighbor.Stone.Id, int64(neighbor.Distance))
			}
		}
	}
}

func addVertexToGraph(grid HexGrid.Grid, graph *dijkstra.Graph, player int) {
	for _, stone := range grid.Stones {
		if stone.Player == player || stone.Player == HexGrid.Empty {
			graph.AddVertex(stone.Id)
		}
	}
}
