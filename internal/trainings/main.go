package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikodem-kirsz/wild-workouts/internal/common/server"
	"github.com/nikodem-kirsz/wild-workouts/internal/trainings/ports"
)

func main() {
	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(), router)
	})
}
