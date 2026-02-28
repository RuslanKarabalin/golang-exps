package store

import (
	"serv/internal/model"
	"sync"
)

type Server struct {
	NextId int
	Somes  map[int]model.SomeType
	Mtx    sync.RWMutex
}
