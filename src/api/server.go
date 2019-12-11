package api

import (
	"api/handler"
	"net/http"
)

// Start the server
// By default, listen on 8080
func Start() {

	http.HandleFunc("/games/is-won", handler.IsWonHandler)

	http.ListenAndServe(":8080", nil)
}
