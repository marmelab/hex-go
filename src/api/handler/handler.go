package handler

import (
	"encoding/json"
	"fmt"
	"hex/game"
	"io"
	"net/http"
)

func IsWonHandler(responseWriter http.ResponseWriter, request *http.Request) {
	currentGame := getGameByRequest(request)

	response := fmt.Sprintf(`{"isWon": %t}`, game.IsWinningGame(currentGame))

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set("Content-Type", "application/json")
	io.WriteString(responseWriter, response)
}

func getGameByRequest(request *http.Request) game.Game {
	decoder := json.NewDecoder(request.Body)
	var currentGame game.Game
	err := decoder.Decode(&currentGame)

	if err != nil {
		panic(err)
	}
	return currentGame
}
