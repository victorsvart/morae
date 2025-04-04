package authroute

import (
	"morae/cmd/routing/router"
	"morae/internal/handler/authhandler"
)

type AuthRoutes struct {
  Handlers *authhandler.AuthHandler
}

func (a *AuthRoutes) Register(r *router.Group) {
	auth := r.SubGroup("/auth")
	auth.Post("/login", a.Handlers.Login)
	auth.Post("/register", a.Handlers.Register)
	auth.Post("/logout", a.Handlers.Logout)
}
