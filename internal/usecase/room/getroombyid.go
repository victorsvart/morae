// Package room contains use cases for room management.
package room

import (
	"context"
	"errors"
	"morae/internal/domain/roomdomain"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

// GetRoomByIDUsecase defines the interface for retrieving a room by ID.
type GetRoomByIDUsecase interface {
	Execute(context.Context, string) (roomdomain.Room, error)
}

// GetRoomByID implements the use case to retrieve a room by its ID.
type GetRoomByID struct {
	repo mongodb.RoomRepository
}

// Execute retrieves a room from the repository using its ID.
func (g *GetRoomByID) Execute(ctx context.Context, id string) (roomdomain.Room, error) {
	if id == "" {
		return roomdomain.Room{}, ErrInvalidID
	}

	doc, err := g.repo.GetRoomByID(ctx, id)
	if doc == nil {
		return roomdomain.Room{}, errors.New("no document found")
	}

	if err != nil {
		return roomdomain.Room{}, err
	}

	return roommapper.ToDomain(doc), nil
}

// ErrInvalidID is returned when an ID is empty or invalid.
var ErrInvalidID = errors.New("id cannot be empty")
