package store

import (
	"context"
	"log/slog"
	"serv/internal/model"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Conn *pgxpool.Pool
	Mtx  sync.RWMutex
}

func (s *Store) GetAllSomes() ([]model.SomeType, error) {
	s.Mtx.Lock()
	defer s.Mtx.Unlock()
	getAllSomes := "select id, name from somes"
	rows, err := s.Conn.Query(context.Background(), getAllSomes)
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

func (s *Store) GetSomeById(id int) (*model.SomeType, error) {
	s.Mtx.Lock()
	defer s.Mtx.Unlock()
	getSomeById := "select id, name from somes where id = $1"
	row := s.Conn.QueryRow(context.Background(), getSomeById, id)
	some := model.SomeType{}
	err := row.Scan(&some.Id, &some.Name)
	if err != nil {
		slog.Warn("Can't scan some", slog.Any("error", err))
		return nil, err
	}
	return &some, nil
}

func (s *Store) InsertSome(some model.CreateSomeRequest) (int, error) {
	id := -1
	s.Mtx.RLock()
	defer s.Mtx.RUnlock()
	insertSome := "insert into somes(name) values($1) returning id"
	genId := s.Conn.QueryRow(context.Background(), insertSome, some.Name)
	err := genId.Scan(&id)
	if err != nil {
		slog.Warn("Can't scan some", slog.Any("error", err))
		return id, err
	}
	return id, nil
}
