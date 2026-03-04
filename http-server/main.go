package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"serv/internal/config"
	"serv/internal/handler"
	"serv/internal/store"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	pgUrl := cfg.GetPgUrl()

	slog.Info("PgUrl:", slog.Any("pgUrl", pgUrl))

	conn, err := pgxpool.New(ctx, pgUrl)
	if err != nil {
		slog.Error("Cannot connect to PostgreSQL", slog.Any("error", err))
		os.Exit(1)
	}
	defer conn.Close()

	// bad migration for start
	createTableSomes := `
	create table if not exists somes(
		id int generated always as identity primary key,
		name text not null
	)
	`
	_, err = conn.Exec(ctx, createTableSomes)
	if err != nil {
		slog.Error("Cannot create somes table", slog.Any("error", err))
	}

	insertSome := `
	insert into somes(name)
	values('Alice')
	`

	_, err = conn.Exec(ctx, insertSome)
	if err != nil {
		slog.Error("Cannot insert into somes table", slog.Any("error", err))
	}

	mux := http.NewServeMux()

	store := &store.Store{
		Conn: conn,
	}

	mux.HandleFunc("POST /somes", handler.PostSomeHandler(store))
	mux.HandleFunc("GET /somes", handler.GetAllSomeHandler(store))
	mux.HandleFunc("GET /somes/{id}", handler.GetSomeByIdHandler(store))
	// mux.HandleFunc("PUT /somes/{id}", handler.PutSomeHandler(store))
	// mux.HandleFunc("DELETE /somes/{id}", handler.DeleteSomeByIdHandler(store))

	addr := cfg.GetAddr()
	slog.Info("Server listening on", slog.String("address", addr))

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Server starting failed", slog.Any("error", err))
		os.Exit(1)
	}
}
