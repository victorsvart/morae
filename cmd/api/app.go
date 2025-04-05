package main

import (
	"log"
	"morae/cmd/config"
	"morae/cmd/migrate/seed"
	"morae/cmd/routing/factory"
	"morae/internal/db"
	"morae/internal/handler"
	"morae/internal/store/mongodb"
	"morae/internal/store/postgres"
	"net/http"
	"time"
)

// App represents the main application structure with its dependencies.
type App struct {
	config       *config.Config
	storage      *postgres.Storage
	mongostorage *mongodb.MongoStorage
	handlers     *handler.Handlers
	routeFactory *factory.RouteFactory
}

// NewApp initializes and returns a new App instance.
func NewApp() *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.storage = app.setupPostgres()
	app.mongostorage = app.setupMongoDb()
	app.handlers = handler.NewHandlers(app.storage, app.mongostorage)
	app.routeFactory = factory.NewRouteFactory(app.handlers)

	return app
}

// setupPostgres initializes and connects to the Postgres database.
func (a *App) setupPostgres() *postgres.Storage {
	db, err := db.New(
		a.config.Dsn,
		a.config.MaxIdleTime,
		a.config.MaxOpenConns,
		a.config.MaxIdleConns,
	)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to Postgres database!")
	seed.Postgres(db)

	return postgres.NewPostgresStorage(db)
}

// setupMongoDb initializes and connects to the MongoDB database.
func (a *App) setupMongoDb() *mongodb.MongoStorage {
	db := db.ConnectMongo(a.config.MongoDsn)

	log.Println("Connected to Mongo database!")
	seed.MongoDb(db)

	return mongodb.NewMongoStorage(db)
}

// mount initializes and returns the HTTP server and application config.
func mount() (*http.Server, *config.Config) {
	app := NewApp()

	svr := &http.Server{
		Addr:         app.config.Port,
		Handler:      app.routeFactory.Router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return svr, app.config
}
