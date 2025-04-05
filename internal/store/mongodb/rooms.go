package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomDocument struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	OwnerId      uint64             `bson:"_ownerId, omitempty"`
	FullAddress string             `bson:"fullAddress"`
	Street      string             `bson:"street"`
	Number      uint16             `bson:"number"`
	District    string             `bson:"district"`
	State       string             `bson:"state"`
}

type RoomRepository interface {
	GetRoomById(context.Context, string) (*RoomDocument, error)
	CreateRoom(context.Context, *RoomDocument) error
}

type RoomStore struct {
	col *mongo.Collection
}

func ReturnErrorInCollection(operation string, err error) error {
	return errors.New(fmt.Sprintf("%s error in collection Rooms: %v", operation, err))
}

func (r *RoomStore) GetRoomById(ctx context.Context, id string) (*RoomDocument, error) {
  objectId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return nil, err
  }

	filter := bson.M{"_id": objectId}

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

