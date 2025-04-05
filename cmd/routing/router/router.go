// Package router provides routing logic with support for middleware, route grouping,
// and standard HTTP error handling like NotFound and MethodNotAllowed.
package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Route represents a single API route with an HTTP method, pattern, and associated handler.
type Route struct {
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

// RouteOptions defines optional settings for route registration.
type RouteOptions struct {
	MiddlewareExclude []string
}

// Router is a custom HTTP router that manages routes, middleware, and error handlers.
type Router struct {
	*http.ServeMux
	RouteHandler

	routes     map[string][]*Route
	notFound   http.HandlerFunc
	notAllowed http.HandlerFunc
}

// MiddlewareFunc defines the base type for middleware functions.
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Middleware wraps a middleware function with an associated name.
type Middleware struct {
	Name string
	Exec MiddlewareFunc
}

// NewMiddleware creates a new named middleware instance.
func NewMiddleware(name string, f MiddlewareFunc) *Middleware {
	return &Middleware{
		Name: name,
		Exec: f,
	}
}

// NewRouter creates a new Router instance with default configuration.
func NewRouter() *Router {
	router := &Router{
		ServeMux:   http.NewServeMux(),
		routes:     make(map[string][]*Route),
		notFound:   http.NotFoundHandler().ServeHTTP,
		notAllowed: NotAllowedHandler().ServeHTTP,
	}
	router.RouteHandler = RouteHandler{router: router, prefix: ""}
	return router
}

// NotAllowed returns a JSON response with a 405 Method Not Allowed status.
func NotAllowed(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	errorResponse := map[string]string{
		"error": "Method not allowed",
	}

	_ = json.NewEncoder(w).Encode(errorResponse)
}

// NotAllowedHandler returns an HTTP handler that invokes NotAllowed.
func NotAllowedHandler() http.Handler {
	return http.HandlerFunc(NotAllowed)
}

// Handle registers a new route for a specific HTTP method and pattern.
func (r *Router) Handle(route *Route) {
	for _, mw := range r.middleware {
		route.handlerFunc = mw.Exec(route.handlerFunc)
	}

	if r.routes[route.method] == nil {
		r.routes[route.method] = []*Route{}
	}

	r.routes[route.method] = append(r.routes[route.method], route)

	fullPattern := fmt.Sprintf("%s %s", route.method, route.pattern)
	log.Printf("Registered %s", fullPattern)

	r.ServeMux.Handle(fullPattern, route.handlerFunc)
}

// Group creates a new route group with a common prefix.
func (r *Router) Group(prefix string) *Group {
	group := &Group{
		RouteHandler: RouteHandler{router: r, prefix: prefix},
	}

	log.Printf("Registered group: %s", prefix)
	return group
}
