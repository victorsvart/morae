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

type RoomDocument struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	OwnerId     uint64             `bson:"_ownerId, omitempty"`
	FullAddress string             `bson:"fullAddress"`
	Street      string             `bson:"street"`
	Number      uint16             `bson:"number"`
	District    string             `bson:"district"`
	State       string             `bson:"state"`
}

type RoomRepository interface {
	GetRoomById(context.Context, string) (*RoomDocument, error)
	GetAllRooms(ctx context.Context, page int64, perPage int64) ([]*RoomDocument, error)
	CreateRoom(context.Context, *RoomDocument) error
	UpdateRoom(ctx context.Context, document *RoomDocument) error
}

type RoomStore struct {
	col *mongo.Collection
}

func ReturnErrorInCollection(operation string, err error) error {
	return errors.New(fmt.Sprintf("%s error in collection Rooms: %v", operation, err))
}

func (r *RoomStore) GetAllRooms(ctx context.Context, page int64, perPage int64) ([]*RoomDocument, error) {
	skip := page * perPage

	opts := options.Find().
		SetSkip(skip).
		SetLimit(perPage).
		SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []*RoomDocument
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

func (r *RoomStore) UpdateRoom(ctx context.Context, document *RoomDocument) error {
	filter := bson.M{"_id": document.ID}
	update := bson.M{"$set": document}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := r.col.FindOneAndUpdate(ctx, filter, update, opts).Decode(&document)
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrInvalidRoomId = errors.New("Invalid room id")
)
