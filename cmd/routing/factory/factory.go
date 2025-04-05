// Package factory is responsible for assembling and registering all route groups and middlewares.
package factory

import (
	"morae/cmd/routing/factory/authroute"
	"morae/cmd/routing/factory/roomroute"
	"morae/cmd/routing/factory/userroute"
	"morae/cmd/routing/healthcheck"
	"morae/cmd/routing/middleware"
	"morae/cmd/routing/router"
	"morae/internal/handler"
)

// RouteFactory constructs and holds the routing configuration for the application.
type RouteFactory struct {
	Router   *router.Router
	handlers *handler.Handlers
}

// NewRouteFactory initializes and returns a new RouteFactory instance.
func NewRouteFactory(h *handler.Handlers) *RouteFactory {
	route := &RouteFactory{
		Router:   router.NewRouter(),
		handlers: h,
	}

	route.setupRoutes()
	return route
}

// setupGlobalMiddlewares registers middlewares applied to all routes.
func (a *RouteFactory) setupGlobalMiddlewares() {
	a.Router.Use(
		router.NewMiddleware("LogMiddleware", middleware.LoggingMiddleware),
		router.NewMiddleware("JsonMiddleware", middleware.JSONMiddleware),
	)
}

// setupRoutes registers all route groups and global middlewares.
func (a *RouteFactory) setupRoutes() {
	a.setupGlobalMiddlewares()

	api := a.Router.Group("/v1/api")
	api.Get("/healthcheck", healthcheck.Handler, nil)

	routeGroups := []RouteGroup{
		&authroute.AuthRoutes{Handlers: &a.handlers.Auth},
		&userroute.UserRoutes{Handlers: &a.handlers.User},
		&roomroute.RoomRoutes{Handlers: &a.handlers.Room},
	}

	for _, rg := range routeGroups {
		rg.Register(api)
	}
}
