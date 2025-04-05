package room

import "morae/internal/store/mongodb"

type RoomUsecases struct {
	GetAllRooms GetAllRoomsUsecase
	GetById     GetRoomByIdUsecase
	CreateRoom  CreateRoomInterface
	UpdateRoom  UpdateRoomUsecase
	DeleteRoom  DeleteRoomUsecase
}

func NewRoomUsecases(repo mongodb.RoomRepository) *RoomUsecases {
	return &RoomUsecases{
		GetAllRooms: &GetAllRooms{repo},
		GetById:     &GetRoomById{repo},
		CreateRoom:  &CreateRoom{repo},
		UpdateRoom:  &UpdateRoom{repo},
		DeleteRoom:  &DeleteRoom{repo},
	}
}
