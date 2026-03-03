package store

import (
	"context"
	"log/slog"
	"serv/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllSomes(conn *pgxpool.Pool) ([]model.SomeType, error) {
	getAllSomes := "select id, name from somes"
	rows, err := conn.Query(context.Background(), getAllSomes)
	if err != nil {
		slog.Warn("Can't get all somes", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()
	somes := []model.SomeType{}
	for rows.Next() {
		some := model.SomeType{}
		err := rows.Scan(&some.Id, &some.Name)
		if err != nil {
			slog.Warn("Can't scan some", slog.Any("error", err))
			return nil, err
		}
		somes = append(somes, some)
	}
	return somes, nil
}
