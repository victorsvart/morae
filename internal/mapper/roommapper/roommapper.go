package roommapper

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/store/mongodb"
)

func SetFullAddress(street, district string, number uint16) string {
	return fmt.Sprintf("%s, %d - %s", street, number, district)
}

func ToDomain(document *mongodb.RoomDocument) roomdomain.Room {
	return roomdomain.Room{
		ID:          document.ID.Hex(), // converts ObjectID → string
		OwnerId:      document.OwnerId,
		Street:      document.Street,
		Number:      document.Number,
		District:    document.District,
		FullAddress: SetFullAddress(document.Street, document.District, document.Number),
		State:       document.State,
	}
}

func ToDocument(room *roomdomain.Room) (mongodb.RoomDocument, error) {
	var objID primitive.ObjectID
	var err error

	if room.ID != "" {
		objID, err = primitive.ObjectIDFromHex(room.ID) // convert string → ObjectID
		if err != nil {
			return mongodb.RoomDocument{}, fmt.Errorf("invalid room ID: %w", err)
		}
	} else {
		objID = primitive.NewObjectID() // generate new ObjectID if empty
	}

	return mongodb.RoomDocument{
		ID:          objID,
		OwnerId:      room.OwnerId,
		Street:      room.Street,
		Number:      room.Number,
		District:    room.District,
		FullAddress: SetFullAddress(room.Street, room.District, room.Number),
		State:       room.State,
	}, nil
}

func FromInput(input *roomdto.RoomInput) mongodb.RoomDocument {
	return mongodb.RoomDocument{
		OwnerId:      input.OwnerId,
		Street:      input.Street,
		Number:      input.Number,
		District:    input.District,
		FullAddress: SetFullAddress(input.Street, input.District, input.Number),
		State:       input.State,
	}
}
