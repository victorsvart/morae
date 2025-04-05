// Package authroute sets up HTTP routes related to authentication.
package authroute

import (
	"morae/cmd/routing/router"
	"morae/internal/handler/authhandler"
)

// AuthRoutes defines the routes related to authentication and holds the relevant handlers.
type AuthRoutes struct {
	Handlers *authhandler.AuthHandler
}

// Register sets up authentication-related routes under the /auth group.
func (a *AuthRoutes) Register(r *router.Group) {
	auth := r.SubGroup("/auth")
	auth.Post("/login", a.Handlers.Login, nil)
	auth.Post("/register", a.Handlers.Register, nil)
	auth.Post("/logout", a.Handlers.Logout, nil)
}
