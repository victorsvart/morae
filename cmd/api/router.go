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
	routes     map[string][]*Route
	middleware []Middleware
	notFound   http.HandlerFunc
	notAllowed http.HandlerFunc
	mu         sync.RWMutex
}

type Middleware struct {
	name      string
	execution MiddlewareFunc
}

// Base type for middleware application
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func NewMiddleware(name string, f MiddlewareFunc) *Middleware {
	return &Middleware{
		name:      name,
		execution: f,
	}
}

// Creates a new Router instance with default configurations
func newRouter() *Router {
	return &Router{
		ServeMux:   http.NewServeMux(),
		routes:     make(map[string][]*Route),
		notFound:   http.NotFoundHandler().ServeHTTP,
		notAllowed: NotAllowedHandler().ServeHTTP,
	}
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
		handler = mw.execution(handler)
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

// Registers middleware functions to be applied to all routes
func (r *Router) Use(mid Middleware) {
	r.middleware = append(r.middleware, mid)
	log.Printf("Registered global middlewares: %v", mid.name)
}

// Creates a new route group with a common prefix | Ex. /v1/api is a group. /v1/api/users is a subgroup
func (r *Router) Group(prefix string) *Group {
	group := &Group{
		router: r,
		prefix: prefix,
	}

	log.Printf("Registered group: %s", prefix)
	return group
}

// Registers a handler for HTTP GET requests
func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}

// Registers a handler for HTTP POST requests
func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, pattern, handler)
}

// Registers a handler for HTTP PUT requests
func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPut, pattern, handler)
}

// Registers a handler for HTTP DELETE requests
func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}

// Implements ServeMux ServeHTTP function. Processes HTTP requests using registerd routes
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.ServeMux.ServeHTTP(w, req)
}
