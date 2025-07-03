package handlers

import (
	"learng/httpgordle/internal/api"
	"learng/httpgordle/internal/handlers/newgame"
	"net/http"
)

// Mux creates a multiplexer with all the endpoints for our service
func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}
