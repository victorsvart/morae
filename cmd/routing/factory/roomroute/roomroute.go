// Package roomroute defines the routing logic for room-related endpoints.
package roomroute

import (
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler/roomhandler"
)

// RoomRoutes sets up the HTTP routes for room operations.
type RoomRoutes struct {
	Handlers *roomhandler.RoomHandler
}

// Register sets up the room-related routes under the provided router group.
func (rr *RoomRoutes) Register(r *router.Group) {
	room := r.SubGroup(
		"/rooms",
		router.NewMiddleware("AuthMiddleware", middleware.AuthMiddleware),
	)

	room.Get("/", rr.Handlers.GetAllRooms, &router.RouteOptions{
		MiddlewareExclude: []string{"AuthMiddleware"},
	})

	room.Get("/{id}", rr.Handlers.GetRoomUserID, nil)
	room.Post("/", rr.Handlers.CreateRoom, nil)
	room.Put("/", rr.Handlers.UpdateRoom, nil)
	room.Delete("/{id}", rr.Handlers.DeleteRoom, nil)
}
