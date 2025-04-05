// Package mongodb provides MongoDB-based storage implementations for repositories.
package mongodb

import "go.mongodb.org/mongo-driver/mongo"

// MongoStorage holds all MongoDB-backed repositories.
type MongoStorage struct {
	Rooms RoomRepository
}

// NewMongoStorage initializes a new MongoStorage with the provided MongoDB database.
func NewMongoStorage(db *mongo.Database) *MongoStorage {
	return &MongoStorage{
		Rooms: &RoomStore{col: db.Collection("rooms")},
	}
}
