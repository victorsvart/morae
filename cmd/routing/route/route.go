package route

import (
	"encoding/json"
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler"
	"net/http"
)

type Route struct {
	Router   *router.Router
	handlers *handler.Handlers
}

func NewRoute(h *handler.Handlers) *Route {
	route := &Route{
		Router:   router.NewRouter(),
		handlers: h,
	}

	route.setupRoutes()
	return route
}

// Registers middlewares
func (a *Route) setupGlobalMiddlewares() {
	a.Router.Use(
		router.Middleware{Name: "LogMiddleware", Exec: middleware.LoggingMiddleware},
		router.Middleware{Name: "JsonMiddleware", Exec: middleware.JsonMiddleware},
	)
}

func (a *Route) setupRoutes() {
	api := a.Router.Group("/v1/api")
	api.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API is up and running",
		})
	})

	// dumb shit. needs a better way to deal with groups and subgroups in the router
	auth := api.SubGroup("/auth")
	auth.Post("/login", a.handlers.Auth.Login)
	auth.Post("/register", a.handlers.Auth.Register)
	auth.Post("/logout", a.handlers.Auth.Logout)

	users := api.SubGroup("/users")
	users.Use(router.Middleware{Name: "AuthMiddleware", Exec: middleware.AuthMiddleware})
	users.Get("/{id}", a.handlers.User.GetUserById)
	users.Post("/", a.handlers.User.CreateUser)
	users.Put("/", a.handlers.User.UpdateUser)
	users.Delete("/{id}", a.handlers.User.DeleteUser)
}
