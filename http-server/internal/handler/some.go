package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"serv/internal/model"
	"serv/internal/store"
	"strconv"
)

func GetAllSomeHandler(s *store.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'getAllSomeHandler'")
		w.Header().Set("Content-Type", "application/json")
		s.Mtx.RLock()
		defer s.Mtx.RUnlock()
		if err := json.NewEncoder(w).Encode(s.Somes); err != nil {
			slog.Error("Can't encode somes", slog.Any("error", err))
		}
	}
}

func GetSomeByIdHandler(s *store.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'getSomeByIdHandler'")

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error("Can't parse id", slog.Any("error", err))
			http.Error(w, "Can't parse id", http.StatusBadRequest)
			return
		}

		s.Mtx.RLock()
		defer s.Mtx.RUnlock()
		some, exists := s.Somes[id]
		if !exists {
			slog.Error("Can't find id")
			http.Error(w, "Can't find id", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(some); err != nil {
			slog.Error("Can't encode some", slog.Any("error", err))
		}
	}
}

func PostSomeHandler(s *store.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'postSomeHandler'")

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var tmp model.SomeType

		if err := dec.Decode(&tmp); err != nil {
			slog.Error("Can't decode body", slog.Any("error", err))
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		s.Mtx.Lock()
		defer s.Mtx.Unlock()
		s.Somes[s.NextId] = tmp
		s.NextId++

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]int{"id": s.NextId - 1})
	}
}

func PutSomeHandler(s *store.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'putSomeHandler'")

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error("Can't parse id", slog.Any("error", err))
			http.Error(w, "Can't parse id", http.StatusBadRequest)
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var tmp model.SomeType

		if err := dec.Decode(&tmp); err != nil {
			slog.Error("Can't decode body", slog.Any("error", err))
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		s.Mtx.Lock()
		defer s.Mtx.Unlock()
		if _, exists := s.Somes[id]; !exists {
			slog.Error("Can't find id")
			http.Error(w, "Can't find id", http.StatusNotFound)
			return
		}
		s.Somes[id] = tmp
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteSomeByIdHandler(s *store.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Call 'deleteSomeByIdHandler'")

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			slog.Error("Can't parse id", slog.Any("error", err))
			http.Error(w, "Can't parse id", http.StatusBadRequest)
			return
		}

		s.Mtx.Lock()
		defer s.Mtx.Unlock()
		if _, exists := s.Somes[id]; !exists {
			slog.Error("Can't find id")
			http.Error(w, "Can't find id", http.StatusNotFound)
			return
		}
		delete(s.Somes, id)
		w.WriteHeader(http.StatusNoContent)
	}
}
