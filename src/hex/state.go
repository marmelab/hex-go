package main

import (
	"github.com/RyanCarrier/dijkstra"
)

func CanWinInOneTurn(bestPath dijkstra.BestPath) bool {
	return bestPath.Distance == 1
}

func isWon(bestPath dijkstra.BestPath) bool {
	return bestPath.Distance == 0
}

func IsClosestAsOpponent(bestPath dijkstra.BestPath, bestPathOpponent dijkstra.BestPath) bool {
	return bestPath.Distance < bestPathOpponent.Distance
}
