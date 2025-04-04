package main

import (
	"encoding/json"
	"morae/internal/db"
	"morae/internal/handler"
	"morae/internal/store/postgres"
	"log"
	"net/http"
	"time"
)

type App struct {
	router   *Router
	config   *Config
	storage  *postgres.PostgresStorage
	handlers *handler.Handlers
}

// Initializes and returns a new App instance
func NewApp() *App {
	cfg := newConfig()
	r := newRouter()

	app := &App{
		router: r,
		config: cfg,
	}

	app.storage = app.setupDatabase()
	app.handlers = handler.NewHandlers(app.storage)

	app.setupGlobalMiddlewares()
	app.setupRoutes()

	return app
}

// Registers middlewares
func (a *App) setupGlobalMiddlewares() {
	a.router.Use(
		Middleware{name: "LogMiddleware", exec: loggingMiddleware},
		Middleware{name: "JsonMiddleware", exec: jsonMiddleware},
	)
}

// Registers API routes
func (a *App) setupRoutes() {
	api := a.router.Group("/v1/api")
	api.Get("/healthcheck", a.healthCheckHandler)

	// dumb shit. needs a better way to deal with groups and subgroups in the router
	auth := api.SubGroup("/auth")
	auth.Post("/login", a.handlers.Auth.Login)
	auth.Post("/register", a.handlers.Auth.Register)

	users := api.SubGroup("/users")
	users.Use(Middleware{name: "AuthMiddleware", exec: authMiddleware})
	users.Get("/{id}", a.handlers.User.GetUserById)
	users.Post("/", a.handlers.User.CreateUser)
	users.Put("/", a.handlers.User.UpdateUser)
	users.Delete("/{id}", a.handlers.User.DeleteUser)
}

// Initializes and connects to the database
func (a *App) setupDatabase() *postgres.PostgresStorage {
	db, err := db.New(
		a.config.dsn,
		a.config.maxIdleTime,
		a.config.maxOpenConns,
		a.config.maxIdleConns,
	)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database!")
	return postgres.NewPostgresStorage(db)
}

// Handles API health checks
func (a *App) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "API is up and running",
	})
}

// Initializes the server
func mount() (*http.Server, *Config) {
	app := NewApp()

	svr := &http.Server{
		Addr:         app.config.port,
		Handler:      app.router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return svr, app.config
}
