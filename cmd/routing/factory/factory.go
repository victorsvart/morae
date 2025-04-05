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

type RouteFactory struct {
	Router   *router.Router
	handlers *handler.Handlers
}

func NewRouteFactory(h *handler.Handlers) *RouteFactory {
	route := &RouteFactory{
		Router:   router.NewRouter(),
		handlers: h,
	}

	route.setupRoutes()
	return route
}

// Registers middlewares
func (a *RouteFactory) setupGlobalMiddlewares() {
	a.Router.Use(
    router.NewMiddleware("LogMiddleware", middleware.LoggingMiddleware),
    router.NewMiddleware("JsonMiddleware", middleware.JsonMiddleware),
	)
}

func (a *RouteFactory) setupRoutes() {
	a.setupGlobalMiddlewares()
	api := a.Router.Group("/v1/api")
	api.Get("/healthcheck", healthcheck.HealthCheckHandler)

	routeGroups := []RouteGroup {
		&authroute.AuthRoutes{Handlers: &a.handlers.Auth},
    &userroute.UserRoutes{Handlers: &a.handlers.User},
    &roomroute.RoomRoutes{Handlers: &a.handlers.Room},
	}

	for _, rg := range routeGroups {
		rg.Register(api)
	}
}
