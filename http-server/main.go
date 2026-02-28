package main

import (
	"log/slog"
	"net/http"
	"serv/internal/handler"
	"serv/internal/model"
	"serv/internal/store"
)

func main() {
	mux := http.NewServeMux()

	server := &store.Server{
		Somes: make(map[int]model.SomeType),
	}

	mux.HandleFunc("GET /somes", handler.GetAllSomeHandler(server))
	mux.HandleFunc("GET /somes/{id}", handler.GetSomeByIdHandler(server))
	mux.HandleFunc("POST /somes", handler.PostSomeHandler(server))
	mux.HandleFunc("PUT /somes/{id}", handler.PutSomeHandler(server))
	mux.HandleFunc("DELETE /somes/{id}", handler.DeleteSomeByIdHandler(server))

	slog.Info("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
	}
}
