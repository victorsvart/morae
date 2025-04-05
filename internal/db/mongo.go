package db

import (
	"context"
	"log"
	"morae/internal/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo returns a Mongo client or database handle
func ConnectMongo(uri string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Panic("Could not connect to MongoDB:", err)
	}

	return client.Database(env.GetString("MONGO_DB", "moraedb"))
}
