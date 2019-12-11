package handler

import (
	"encoding/json"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"hex/game"
	Graph "hex/graph"
	"hex/grid"
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

	graph := &dijkstra.Graph{}
	var endVertexId int
	if currentGame.Player == grid.Player1 {
		graph, endVertexId = Graph.BootstrapGraphPlayer1(currentGame.Matrix)


	} else if currentGame.Player == grid.Player2 {
		graph, endVertexId = Graph.BootstrapGraphPlayer2(currentGame.Matrix)
	}

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)
	response := fmt.Sprintf(`{"isWon": %t}`, state.IsWon(bestPath))
	_, _ = io.WriteString(w, response)

}
