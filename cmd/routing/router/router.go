package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Represents a single API route with a pattern and asssociated handler function
type Route struct {
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type RouteOptions struct {
	MiddlewareExclude []string
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
	Name string
	Exec MiddlewareFunc
}

// Base type for middleware application
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func NewMiddleware(name string, f MiddlewareFunc) *Middleware {
	return &Middleware{
		Name: name,
		Exec: f,
	}
}

// Creates a new Router instance with default configurations
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
func (r *Router) Handle(route *Route) {
	for _, mw := range r.middleware {
		route.handlerFunc = mw.Exec(route.handlerFunc)
	}

	if r.routes[route.method] == nil {
		r.routes[route.method] = []*Route{}
	}

	r.routes[route.method] = append(r.routes[route.method], route)

	// method + pattern | Ex. GET /v1/api/users
	fullPattern := fmt.Sprintf("%s %s", route.method, route.pattern)
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
