package room

import (
	"context"
	"errors"
	"morae/internal/domain/roomdomain"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

type GetRoomByIdUsecase interface {
	Execute(context.Context, string) (roomdomain.Room, error)
}

type GetRoomById struct {
	repo mongodb.RoomRepository
}

func (g *GetRoomById) Execute(ctx context.Context, id string) (roomdomain.Room, error) {
	if id == "" {
		return roomdomain.Room{}, ErrInvalidId
	}

	doc, err := g.repo.GetRoomById(ctx, id)
	if doc == nil {
		return roomdomain.Room{}, errors.New("No document found")
	}

	if err != nil {
		return roomdomain.Room{}, err
	}

	return roommapper.ToDomain(doc), nil
}

var (
	ErrInvalidId = errors.New("Id cannot be empty")
)
