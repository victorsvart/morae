// Package room provides use cases for managing room entities.
package room

import (
	"context"
	"errors"
	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

// GetAllRoomsUsecase defines the contract for retrieving a list of rooms with pagination.
type GetAllRoomsUsecase interface {
	Execute(context.Context, *roomdto.GetRoomPaged) ([]*roomdomain.Room, error)
}

// GetAllRooms is the concrete implementation of GetAllRoomsUsecase.
type GetAllRooms struct {
	repo mongodb.RoomRepository
}

// Execute retrieves all rooms based on pagination parameters.
func (g *GetAllRooms) Execute(ctx context.Context, input *roomdto.GetRoomPaged) ([]*roomdomain.Room, error) {
	docs, err := g.repo.GetAllRooms(ctx, input.Page, input.PerPage)
	if docs == nil {
		return nil, errors.New("no documents found")
	}

	if err != nil {
		return nil, err
	}

	return roommapper.ToDomainSlice(docs), nil
}
