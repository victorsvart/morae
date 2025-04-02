package main

import (
	"database/sql"
	"encoding/json"
	"golangproj/internal/db"
	"golangproj/internal/handler"
	"golangproj/internal/store/postgres"
	"log"
	"net/http"
	"time"
)

type App struct {
	router   *Router
	config   *Config
	handlers *handler.Handlers
}

func mount() (*http.Server, *Config) {
	cfg := newConfig()
	r := newRouter()

	app := &App{
		router: r,
		config: cfg,
	}

	storage := app.setupDatabase()
	app.handlers = app.setupHandlers(storage)

	r.Use(Middleware{name: "LogMiddleware", execution: loggingMiddleware})
	r.Use(Middleware{name: "JsonMiddleware", execution: jsonMiddleware})

	api := r.Group("/v1/api")
	api.Get("/health", app.healthCheckHandler)
	api.Get("/users", app.handlers.User.CreateUser)

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

func (a *App) setupHandlers(storage *postgres.PostgresStorage) *handler.Handlers {
	handlers := handler.NewHandlers(storage)
	return handlers
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
