package main

import (
	"database/sql"
	"encoding/json"
	"golangproj/internal/db"
	"golangproj/internal/store/postgres"
	"log"
	"net/http"
	"time"
)

type App struct {
	router *Router
	config *Config
	store  *postgres.PostgresStorage
}

func mount() (*http.Server, *Config) {
	cfg := newConfig()
	r := newRouter()

	app := &App{
		router: r,
		config: cfg,
	}
	app.store = app.setupDatabase()

	r.Use(Middleware{name: "LogMiddleware", execution: loggingMiddleware})
	r.Use(Middleware{name: "JsonMiddleware", execution: jsonMiddleware})

	api := r.Group("/v1/api")
	api.Get("/health", app.healthCheckHandler)

	svr := &http.Server{
		Addr:         app.config.port,
		Handler:      app.router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return svr, cfg
}

func setupStorage(db *sql.DB) *postgres.PostgresStorage {
	postgresCon := postgres.NewPostgresStorage(db)
	return postgresCon

}

func (a *App) setupDatabase() *postgres.PostgresStorage {
	db, err := db.New(
		a.config.dsn,
		a.config.maxIdleTime,
		a.config.maxOpenConns,
		a.config.maxIdleConns,
	)

	if err != nil {
		log.Panic("Error connecting to database!")
	}

	log.Println("Connected to database!")
	return setupStorage(db)
}

func (a *App) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "ok", "message": "API is up and running"}
	json.NewEncoder(w).Encode(response)
}
