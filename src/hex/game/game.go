package game

import (
	"hex/graph"
	"hex/grid"
	"hex/state"
)

type Game struct {
	Array  []int
	Player int
}

// This function determines if a game is won based on a game
func IsWinningGame(game Game) bool {
	Grid := grid.GetGridFromArray(game.Array)
	bestPath, _ := graph.GetBestPathForGraph(Grid, game.Player)
	return state.IsWon(bestPath)
}

// This function determines which stone should be played at the next time
func GetAdvice(game Game) int {
	Grid := grid.GetGridFromArray(game.Array)
	bestPath, Graph := graph.GetBestPathForGraph(Grid, game.Player)
	return state.AdviceForNextMove(bestPath, Graph)
}
