package main

import (
	"github.com/RyanCarrier/dijkstra"
)

func CanWinInOneTurn(bestPath dijkstra.BestPath) bool {
	return bestPath.Distance == 1
}

func IsWon(bestPath dijkstra.BestPath) bool {
	return bestPath.Distance == 0
}

func IsClosestAsOpponent(bestPath dijkstra.BestPath, bestPathOpponent dijkstra.BestPath) bool {
	return bestPath.Distance < bestPathOpponent.Distance
}

func AdviceForNextMove(bestPath dijkstra.BestPath, graph *dijkstra.Graph) int {

	for i, id := range bestPath.Path {
		vertex, _ := graph.GetVertex(id)
		next := bestPath.Path[i+1]
		distance, _ := vertex.GetArc(next)

		if distance > 0 {
			return next
		}

	}

	return -1
}
