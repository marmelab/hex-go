package game

import (
	"hex/graph"
	"hex/grid"
	"hex/state"
)

type Game struct {
	Array []int
	Player int
}


// This function determines if a game is won based on a game
func IsWinningGame(game Game) bool {

	Grid := grid.GetGridFromArray(game.Array)

	endVertexId := graph.GetEndVertexId(Grid.Width)
	Graph := graph.BuildGraphForPlayer(Grid, game.Player)

	bestPath, _ := Graph.Shortest(graph.StartVertexId, endVertexId)

	return state.IsWon(bestPath)
}
