package main

import (
	"log"
	"morae/cmd/config"
	"morae/cmd/routing/route"
	"morae/internal/db"
	"morae/internal/handler"
	"morae/internal/store/postgres"
	"net/http"
	"time"
)

type App struct {
	config   *config.Config
	storage  *postgres.PostgresStorage
	handlers *handler.Handlers
	route   *route.Route
}

// Initializes and returns a new App instance
func NewApp() *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.storage = app.setupDatabase()
	app.handlers = handler.NewHandlers(app.storage)
  app.route = route.NewRoute(app.handlers)
	return app
}

// Initializes and connects to the database
func (a *App) setupDatabase() *postgres.PostgresStorage {
	db, err := db.New(
		a.config.Dsn,
		a.config.MaxIdleTime,
		a.config.MaxOpenConns,
		a.config.MaxIdleConns,
	)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database!")
	return postgres.NewPostgresStorage(db)
}

// Initializes the server
func mount() (*http.Server, *config.Config) {
	app := NewApp()

	svr := &http.Server{
		Addr:         app.config.Port,
		Handler:      app.route.Router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return svr, app.config
}
