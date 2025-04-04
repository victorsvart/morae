package userroute

import (
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler/userhandler"
)

type UserRoutes struct {
  Handlers *userhandler.UserHandler
}

func (u *UserRoutes) Register(r *router.Group) {
	users := r.SubGroup(
		"/users",
    router.NewMiddleware("AuthMiddleware", middleware.AuthMiddleware),
	)

	users.Get("/{id}", u.Handlers.GetUserById)
	users.Post("/", u.Handlers.CreateUser)
	users.Put("/", u.Handlers.UpdateUser)
	users.Delete("/{id}", u.Handlers.DeleteUser)
}
