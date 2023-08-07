package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	server "github.com/nikodemkirsz/wild-workouts/internal/common"
	ports "github.com/nikodemkirsz/wild-workouts/internal/trainings"
)

func main() {
	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(), router)
	})
}
