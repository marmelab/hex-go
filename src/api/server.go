package api

import (
	"api/handler"
	"log"
	"net/http"
)

// Start the server
// By default, listen on 8080
func Start() {

	http.HandleFunc("/games/advices", handler.GetAdviceHandler)
	http.HandleFunc("/games/is-won", handler.IsWonHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
