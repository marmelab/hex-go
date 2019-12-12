package handler

import (
	"encoding/json"
	"fmt"
	"hex/game"
	"io"
	"net/http"
)

func IsWonHandler(responseWriter http.ResponseWriter, request *http.Request) {

	setupResponse(&responseWriter)
	if (*request).Method == "OPTIONS" {
		return
	}

	currentGame := getGameByRequest(request)
	response := fmt.Sprintf(`{"isWon": %t}`, game.IsWinningGame(currentGame))

	responseWriter.WriteHeader(http.StatusOK)
	io.WriteString(responseWriter, response)
}

func GetAdviceHandler(responseWriter http.ResponseWriter, request *http.Request) {

	setupResponse(&responseWriter)
	if (*request).Method == "OPTIONS" {
		return
	}

	currentGame := getGameByRequest(request)
	response := fmt.Sprintf(`{"advice": %d}`, game.GetAdvice(currentGame))

	responseWriter.WriteHeader(http.StatusOK)
	io.WriteString(responseWriter, response)
}

func setupResponse(writer *http.ResponseWriter) {
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	(*writer).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*writer).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, content-type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
