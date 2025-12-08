package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nambuitechx/go-monitor/backend/configs"
	healthlogs "github.com/nambuitechx/go-monitor/backend/health_logs"
)

func NewRouter(envConfig *configs.EnvConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Init

	// // Scylla connection
	// db := configs.NewScyllaConnection(envConfig)

	// Postgres connection
	db := configs.NewPostgresConnection(envConfig)

	// Health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "healthy",
		})
	})

	// health-logs route
	healthLogsRepo := healthlogs.NewHealthLogRepository(db)
	healthLogsService := healthlogs.NewHealthLogService(healthLogsRepo)
	healthLogsRouter := healthlogs.NewHealthLogRouter(healthLogsService)

	r.Mount("/health-logs", healthLogsRouter)

	return r
}
