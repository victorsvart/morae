package room

import "morae/internal/store/mongodb"

type RoomUsecases struct {
	GetAllRooms GetAllRoomsUsecase
	GetById     GetRoomByIdUsecase
	CreateRoom  CreateRoomInterface
}

func NewRoomUsecases(repo mongodb.RoomRepository) *RoomUsecases {
	return &RoomUsecases{
    GetAllRooms: &GetAllRooms{repo},
		GetById:    &GetRoomById{repo},
		CreateRoom: &CreateRoom{repo},
	}
}
