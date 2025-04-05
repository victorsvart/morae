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
  authMiddlewareName := "AuthMiddleware"
	room := r.SubGroup(
		"/rooms",
		router.NewMiddleware(authMiddlewareName, middleware.AuthMiddleware),
	)

	room.Get("/", rr.Handlers.GetAllRooms, &authMiddlewareName)
	room.Get("/{id}", rr.Handlers.GetRoomUserId)
	room.Post("/", rr.Handlers.CreateRoom)
	room.Put("/", rr.Handlers.UpdateRoom)
}
