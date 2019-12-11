package handler

import (
	"encoding/json"
	"fmt"
	"hex/game"
	Graph "hex/graph"
	HexGrid "hex/grid"
	"hex/state"
	"io"
	"net/http"
)

func IsWonHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var currentGame game.Game
	err := decoder.Decode(&currentGame)

	if err != nil {
		panic(err)
	}

	grid := HexGrid.GetGridFromMatrix(currentGame.Matrix)

	endVertexId := Graph.GetEndVertexId(grid.Width)
	graph := Graph.BuildGraphForPlayer(grid, currentGame.Player)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)


	response := fmt.Sprintf(`{"isWon": %t}`, state.IsWon(bestPath))
	_, _ = io.WriteString(w, response)

}
