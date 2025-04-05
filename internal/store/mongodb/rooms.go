// Package mongodb provides the MongoDB implementation for room data persistence and access logic.
package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RoomDocument represents a room entity in MongoDB.
type RoomDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID     uint64             `bson:"_ownerId,omitempty"`
	FullAddress string             `bson:"fullAddress"`
	Street      string             `bson:"street"`
	Number      uint16             `bson:"number"`
	District    string             `bson:"district"`
	State       string             `bson:"state"`
}

// RoomRepository defines the behavior for room persistence layer.
type RoomRepository interface {
	GetRoomByID(context.Context, string) (*RoomDocument, error)
	GetAllRooms(ctx context.Context, page int64, perPage int64) ([]*RoomDocument, error)
	CreateRoom(context.Context, *RoomDocument) error
	UpdateRoom(ctx context.Context, document *RoomDocument) error
	DeleteRoom(ctx context.Context, id string) error
}

// RoomStore implements RoomRepository using a MongoDB collection.
type RoomStore struct {
	col *mongo.Collection
}

// ReturnErrorInCollection formats error messages for MongoDB operations.
func ReturnErrorInCollection(operation string, err error) error {
	return fmt.Errorf("%s error in collection Rooms: %w", operation, err)
}

// GetAllRooms returns a paginated list of all room documents sorted by creation time.
func (r *RoomStore) GetAllRooms(ctx context.Context, page int64, perPage int64) (rooms []*RoomDocument, err error) {
	skip := page * perPage

	opts := options.Find().
		SetSkip(skip).
		SetLimit(perPage).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for cursor.Next(ctx) {
		var room RoomDocument
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

// GetRoomByID fetches a room document by its ID.
func (r *RoomStore) GetRoomByID(ctx context.Context, id string) (*RoomDocument, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidRoomID
	}

	filter := bson.M{"_id": objectID}

	var doc RoomDocument
	err = r.col.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, ReturnErrorInCollection("FindOne", err)
	}

	return &doc, nil
}

// CreateRoom inserts a new room document into the collection.
func (r *RoomStore) CreateRoom(ctx context.Context, document *RoomDocument) error {
	if document.ID.IsZero() {
		document.ID = primitive.NewObjectID()
	}

	result, err := r.col.InsertOne(ctx, document)
	if err != nil {
		return ReturnErrorInCollection("InsertOne", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		document.ID = oid
	}

	return nil
}

// UpdateRoom updates an existing room document.
func (r *RoomStore) UpdateRoom(ctx context.Context, document *RoomDocument) error {
	filter := bson.M{"_id": document.ID}
	update := bson.M{"$set": document}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	if err := r.col.FindOneAndUpdate(ctx, filter, update, opts).Decode(&document); err != nil {
		return err
	}

	return nil
}

// DeleteRoom removes a room document by its ID.
func (r *RoomStore) DeleteRoom(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidRoomID
	}

	filter := bson.M{"_id": objectID}
	result := r.col.FindOneAndDelete(ctx, filter)
	return result.Err()
}

// ErrInvalidRoomID is returned when a room ID is not a valid MongoDB ObjectID.
var ErrInvalidRoomID = errors.New("invalid room id")
