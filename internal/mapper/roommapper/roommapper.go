// Package roommapper provides functions to convert between domain, DTO, and database representations of rooms.
package roommapper

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/store/mongodb"
)

// SetFullAddress formats a full address string using street, number, and district.
func SetFullAddress(street, district string, number uint16) string {
	return fmt.Sprintf("%s, %d - %s", street, number, district)
}

// ToDomain converts a RoomDocument to a Room domain entity.
func ToDomain(document *mongodb.RoomDocument) roomdomain.Room {
	return roomdomain.Room{
		ID:          document.ID.Hex(),
		OwnerID:     document.OwnerID,
		Street:      document.Street,
		Number:      document.Number,
		District:    document.District,
		FullAddress: SetFullAddress(document.Street, document.District, document.Number),
		State:       document.State,
	}
}

// ToDomainSlice converts a slice of RoomDocuments to a slice of Room domain entities.
func ToDomainSlice(document []*mongodb.RoomDocument) []*roomdomain.Room {
	slice := make([]*roomdomain.Room, len(document))
	for i, doc := range document {
		slice[i] = &roomdomain.Room{
			ID:       doc.ID.Hex(),
			Street:   doc.Street,
			Number:   doc.Number,
			District: doc.District,
			State:    doc.State,
		}
	}
	return slice
}

// ToDocument converts a Room domain entity to a RoomDocument.
func ToDocument(room *roomdomain.Room) (mongodb.RoomDocument, error) {
	objID, err := toNewObjectID(room.ID)
	if err != nil {
		return mongodb.RoomDocument{}, err
	}

	return mongodb.RoomDocument{
		ID:          objID,
		OwnerID:     room.OwnerID,
		Street:      room.Street,
		Number:      room.Number,
		District:    room.District,
		FullAddress: SetFullAddress(room.Street, room.District, room.Number),
		State:       room.State,
	}, nil
}

// FromInput converts a RoomInput DTO to a RoomDocument.
func FromInput(input *roomdto.RoomInput) mongodb.RoomDocument {
	return mongodb.RoomDocument{
		OwnerID:     input.OwnerID,
		Street:      input.Street,
		Number:      input.Number,
		District:    input.District,
		FullAddress: SetFullAddress(input.Street, input.District, input.Number),
		State:       input.State,
	}
}

// FromDto converts a RoomDto to a RoomDocument.
func FromDto(input *roomdto.RoomDto) (mongodb.RoomDocument, error) {
	objID, err := toObjectID(input.ID)
	if err != nil {
		return mongodb.RoomDocument{}, err
	}

	return mongodb.RoomDocument{
		ID:          objID,
		OwnerID:     input.OwnerID,
		Street:      input.Street,
		Number:      input.Number,
		District:    input.District,
		FullAddress: SetFullAddress(input.Street, input.District, input.Number),
		State:       input.State,
	}, nil
}

// toNewObjectID returns an existing ObjectID from a string, or creates a new one if empty.
func toNewObjectID(plainID string) (primitive.ObjectID, error) {
	if plainID != "" {
		objID, err := primitive.ObjectIDFromHex(plainID)
		if err != nil {
			return primitive.ObjectID{}, fmt.Errorf("invalid room ID: %w", err)
		}
		return objID, nil
	}
	return primitive.NewObjectID(), nil
}

// toObjectID returns an ObjectID from a string if it's not empty.
func toObjectID(plainID string) (primitive.ObjectID, error) {
	if plainID != "" {
		objID, err := primitive.ObjectIDFromHex(plainID)
		if err != nil {
			return primitive.ObjectID{}, fmt.Errorf("invalid room ID: %w", err)
		}
		return objID, nil
	}
	return primitive.ObjectID{}, nil
}
