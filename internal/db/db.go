// Package db provides a PostgreSQL database connection setup utility.
package db

import (
	"context"
	"database/sql"
	"time"

	// Import PostgreSQL driver for side-effects (i.e. to register the driver)
	_ "github.com/lib/pq"
)

// New creates and returns a new PostgreSQL database connection with the given configuration.
// It sets max idle time, max open connections, and max idle connections.
// It also verifies the connection using a timeout-based ping.
func New(dsn, maxIdleTime string, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
