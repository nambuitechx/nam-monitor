package hosts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHostRouter(s *HostService) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", GetAll(s))

	return r
}

func GetAll(s *HostService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		healthLogs, err := s.GetAll()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Get all hosts successfully",
			"data": healthLogs,
		})
	})
}

