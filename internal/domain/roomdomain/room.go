// Package roomdomain contains domain-level representations and business logic for rooms.
package roomdomain

// Room represents the core domain model for a room document
type Room struct {
	ID          string
	OwnerID     uint64
	Street      string
	Number      uint16
	District    string
	FullAddress string
	State       string
}
