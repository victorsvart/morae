package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Represents a single API route with a pattern and asssociated handler function
type Route struct {
	pattern     string
	handlerFunc http.HandlerFunc
}

// Custom HTTP router. Manages router, middleware, and basic http error handling
// such as "NotFound and NotAllowed statuses
type Router struct {
	*http.ServeMux
	RouteHandler
	routes     map[string][]*Route
	notFound   http.HandlerFunc
	notAllowed http.HandlerFunc
	mu         sync.RWMutex
}

type Middleware struct {
	name string
	exec MiddlewareFunc
}

// Base type for middleware application
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func NewMiddleware(name string, f MiddlewareFunc) *Middleware {
	return &Middleware{
		name: name,
		exec: f,
	}
}

// Creates a new Router instance with default configurations
func newRouter() *Router {
	router := &Router{
		ServeMux:   http.NewServeMux(),
		routes:     make(map[string][]*Route),
		notFound:   http.NotFoundHandler().ServeHTTP,
		notAllowed: NotAllowedHandler().ServeHTTP,
	}
	router.RouteHandler = RouteHandler{router: router, prefix: ""}

	return router
}

// Default handler for method-not-allowed cases
func NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	errorResponse := map[string]string{
		"error": "Method not allowed",
	}

	json.NewEncoder(w).Encode(errorResponse)
}

// Returns an HTTP handler that invokes NotAllowed HandleFunc
func NotAllowedHandler() http.Handler {
	return http.HandlerFunc(NotAllowed)
}

// Registers a new route  for a specific HTTP method and pattern
func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	for _, mw := range r.middleware {
		handler = mw.exec(handler)
	}

	route := &Route{
		pattern:     pattern,
		handlerFunc: handler,
	}

	if r.routes[method] == nil {
		r.routes[method] = []*Route{}
	}

	r.routes[method] = append(r.routes[method], route)

	// method + pattern | Ex. GET /v1/api/users
	fullPattern := fmt.Sprintf("%s %s", method, pattern)
	log.Printf("Registered %s", fullPattern)
	r.ServeMux.Handle(fullPattern, route.handlerFunc)
}

// Creates a new route group with a common prefix | Ex. /v1/api is a group. /v1/api/users is a subgroup
func (r *Router) Group(prefix string) *Group {
	group := &Group{
		RouteHandler: RouteHandler{router: r, prefix: prefix},
	}

	log.Printf("Registered group: %s", prefix)
	return group
}
