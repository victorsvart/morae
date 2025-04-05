package router

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
)

// RouteHandler provides HTTP route registration with optional middleware and route prefix support.
type RouteHandler struct {
	router     *Router
	middleware []*Middleware
	prefix     string // Prefix used for all routes in this handler (e.g. group/subgroup). May be empty.
}

// Handle registers a route with a specific HTTP method, path pattern, and optional route-specific options.
func (rh *RouteHandler) Handle(method, pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	fullPattern := rh.prefix
	if !strings.HasSuffix(rh.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern

	// Apply middleware unless excluded in RouteOptions
	for _, mw := range rh.middleware {
		if opts != nil && slices.Contains(opts.MiddlewareExclude, mw.Name) {
			log.Printf("Route %s %s excluded from middleware %s", method, fullPattern, mw.Name)
			continue
		}
		handler = mw.Exec(handler)
	}

	route := Route{
		method:      method,
		pattern:     fullPattern,
		handlerFunc: handler,
	}

	rh.router.Handle(&route)
}

// Use registers middleware to be applied to all routes handled by this RouteHandler.
func (rh *RouteHandler) Use(mid ...*Middleware) {
	for _, m := range mid {
		if m == nil {
			continue
		}
		rh.middleware = append(rh.middleware, m)
		scope := "global"
		if rh.prefix != "" {
			scope = fmt.Sprintf("route %q", rh.prefix)
		}
		log.Printf("Applied middleware %q to %s", m.Name, scope)
	}
}

// Get registers a GET route with the given pattern and handler.
func (rh *RouteHandler) Get(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	rh.Handle(http.MethodGet, pattern, handler, opts)
}

// Post registers a POST route with the given pattern and handler.
func (rh *RouteHandler) Post(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	rh.Handle(http.MethodPost, pattern, handler, opts)
}

// Put registers a PUT route with the given pattern and handler.
func (rh *RouteHandler) Put(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	rh.Handle(http.MethodPut, pattern, handler, opts)
}

// Delete registers a DELETE route with the given pattern and handler.
func (rh *RouteHandler) Delete(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	rh.Handle(http.MethodDelete, pattern, handler, opts)
}
