package main

import (
	"log"
	"morae/cmd/config"
	"morae/cmd/routing/factory"
	"morae/internal/db"
	"morae/internal/handler"
	"morae/internal/store/mongodb"
	"morae/internal/store/postgres"
	"net/http"
	"time"
)

type App struct {
	config   *config.Config
	storage  *postgres.PostgresStorage
  mongostorage *mongodb.MongoStorage
	handlers *handler.Handlers
	routeFactory   *factory.RouteFactory
}

// Initializes and returns a new App instance
func NewApp() *App {
	cfg := config.NewConfig()

	app := &App{
		config: cfg,
	}

	app.storage = app.setupDatabase()
  app.mongostorage = app.setupMongoDb()
	app.handlers = handler.NewHandlers(app.storage, app.mongostorage)
  app.routeFactory = factory.NewRouteFactory(app.handlers)
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

	log.Println("Connected to postgres database!")
	return postgres.NewPostgresStorage(db)
}

func (a *App) setupMongoDb() *mongodb.MongoStorage {
  db := db.ConnectMongo(a.config.MongoDsn)

  log.Println("Connected to mongo database!")
  return mongodb.NewMongoStorage(db)
}

// Initializes the server
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
