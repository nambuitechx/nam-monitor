package healthlogs

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHealthLogRouter(s *HealthLogService) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", GetAll(s))

	return r
}

func GetAll(s *HealthLogService) http.HandlerFunc {
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
			"message": "Get all health logs successfully",
			"data": healthLogs,
		})
	})
}
