package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoStorage struct {
  Rooms RoomRepository
}

func NewMongoStorage(db *mongo.Database) *MongoStorage {
  return &MongoStorage{
    Rooms: &RoomStore{col:db.Collection("rooms")},
  }
}
