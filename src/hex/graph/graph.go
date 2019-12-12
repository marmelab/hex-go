package graph

import (
	"errors"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"hex/grid"
)

const StartVertexId = 0

func GetBestPathForGraph(Grid grid.Grid, player int) (dijkstra.BestPath, *dijkstra.Graph) {

	endVertexId := GetEndVertexId(Grid.Width)
	Graph := BuildGraphForPlayer(Grid, player)

	bestPath, _ := Graph.Shortest(StartVertexId, endVertexId)
	return bestPath, Graph
}

func BuildGraphForPlayer(Grid grid.Grid, player int) *dijkstra.Graph {

	var graph *dijkstra.Graph
	if player == grid.Player1 {
		graph = BuildGraphForPlayer1(Grid)
	} else if player == grid.Player2 {
		graph = BuildGraphForPlayer2(Grid)
	} else {
		panic(errors.New("unable to determine player"))
	}

	return graph
}

func BuildGraphForPlayer1(Grid grid.Grid) *dijkstra.Graph {
	graph := initGraph(Grid.Width)

	addVertexToGraph(Grid, graph, grid.Player1)

	addPlayer1StartArcsToGraph(Grid, graph)
	addArcsToGraph(Grid, graph, grid.Player1)
	addPlayer1EndArcsToGraph(Grid, graph)

	return graph
}

func BuildGraphForPlayer2(Grid grid.Grid) *dijkstra.Graph {
	graph := initGraph(Grid.Width)

	addVertexToGraph(Grid, graph, grid.Player2)

	addPlayer2StartArcsToGraph(Grid, graph)
	addArcsToGraph(Grid, graph, grid.Player2)
	addPlayer2EndArcsToGraph(Grid, graph)

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

func addPlayer1StartArcsToGraph(Grid grid.Grid, graph *dijkstra.Graph) {
	for _, stone := range Grid.Stones[:Grid.Width] {
		distance := getDistance(grid.Player1, stone)
		if nil != graph.AddArc(StartVertexId, stone.Id, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2StartArcsToGraph(Grid grid.Grid, graph *dijkstra.Graph) {
	columnIndex := 1
	for i := range Grid.Stones {
		if i == columnIndex {
			currentStone := Grid.Stones[i-1]
			distance := getDistance(grid.Player2, currentStone)

			if nil != graph.AddArc(StartVertexId, currentStone.Id, distance) {
				fmt.Errorf("error during Arc insertion")
			}

			columnIndex = columnIndex + Grid.Width
		}
	}
}

func addPlayer1EndArcsToGraph(Grid grid.Grid, graph *dijkstra.Graph) {
	EndVertexId := GetEndVertexId(Grid.Width)
	offset := (Grid.Width * Grid.Width) - Grid.Width
	for _, stone := range Grid.Stones[offset:] {
		distance := getDistance(grid.Player1, stone)

		if nil != graph.AddArc(stone.Id, EndVertexId, distance) {
			fmt.Errorf("error during Arc insertion")
		}
	}
}

func addPlayer2EndArcsToGraph(Grid grid.Grid, graph *dijkstra.Graph) {
	EndVertexId := GetEndVertexId(Grid.Width)
	columnIndex := Grid.Width
	for i := range Grid.Stones {

		if i+1 == columnIndex {
			currentStone := Grid.Stones[i]
			distance := getDistance(grid.Player2, currentStone)

			if nil != graph.AddArc(currentStone.Id, EndVertexId, distance) {
				fmt.Errorf("error during Arc insertion")
			}
			columnIndex = columnIndex + Grid.Width
		}
	}
}

func getDistance(player int, stone grid.Stone) int64 {
	if player == stone.Player {
		return grid.DistanceOwned
	}

	return grid.DistanceEmpty
}

func addArcsToGraph(Grid grid.Grid, graph *dijkstra.Graph, player int) {
	for _, stone := range Grid.Stones {
		if stone.Player == player || stone.Player == grid.Empty {
			for _, neighbor := range grid.GetNeighborsForStone(Grid, stone, player) {
				graph.AddArc(stone.Id, neighbor.Stone.Id, int64(neighbor.Distance))
			}
		}
	}
}

func addVertexToGraph(Grid grid.Grid, graph *dijkstra.Graph, player int) {
	for _, stone := range Grid.Stones {
		if stone.Player == player || stone.Player == grid.Empty {
			graph.AddVertex(stone.Id)
		}
	}
}
