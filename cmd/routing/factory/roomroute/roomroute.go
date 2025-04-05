package roomroute

import (
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler/roomhandler"
)

type RoomRoutes struct {
	Handlers *roomhandler.RoomHandler
}

func (rr *RoomRoutes) Register(r *router.Group) {
	room := r.SubGroup(
		"/rooms",
		router.NewMiddleware("AuthMiddleware", middleware.AuthMiddleware),
	)

	room.Get("/{id}", rr.Handlers.GetRoomUserId)
	room.Post("/", rr.Handlers.CreateRoom)
}
