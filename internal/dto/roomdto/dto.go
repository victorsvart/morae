// Package roomdto defines data transfer objects used for room-related API operations.
package roomdto

// RoomDto represents the data returned to clients for a room entity.
type RoomDto struct {
	ID       string `json:"id"`
	OwnerID  uint64 `json:"ownerId"`
	Street   string `json:"street"`
	Number   uint16 `json:"number"`
	District string `json:"district"`
	State    string `json:"state"`
}

// RoomInput represents the expected input data when creating or updating a room.
type RoomInput struct {
	OwnerID  uint64 `json:"ownerId"`
	Street   string `json:"street"`
	Number   uint16 `json:"number"`
	District string `json:"district"`
	State    string `json:"state"`
}

// GetRoomPaged contains pagination parameters for listing rooms.
type GetRoomPaged struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"perPage"`
}
