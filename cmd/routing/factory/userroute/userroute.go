// Package userroute sets up HTTP routes related to user operations.
package userroute

import (
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler/userhandler"
)

// UserRoutes defines the routes related to user operations and holds the relevant handlers.
type UserRoutes struct {
	Handlers *userhandler.UserHandler
}

// Register sets up user-related routes under the /users group with authentication middleware.
func (u *UserRoutes) Register(r *router.Group) {
	users := r.SubGroup(
		"/users",
		router.NewMiddleware("AuthMiddleware", middleware.AuthMiddleware),
	)

	users.Get("/{id}", u.Handlers.GetUserByID, nil)
	users.Post("/", u.Handlers.CreateUser, nil)
	users.Put("/", u.Handlers.UpdateUser, nil)
	users.Delete("/{id}", u.Handlers.DeleteUser, nil)
}
