package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RouteHandler struct {
	router     *Router
	middleware []Middleware
	prefix     string
}

func (rh *RouteHandler) Handle(method, pattern string, handler http.HandlerFunc) {
	fullPattern := rh.prefix
	if !strings.HasSuffix(rh.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern
	for i := range rh.middleware {
		handler = rh.middleware[i].Exec(handler)
	}

	rh.router.Handle(method, fullPattern, handler)
}

func (rh *RouteHandler) Use(mid ...Middleware) {
	rh.middleware = append(rh.middleware, mid...)

	scope := "global"
	if rh.prefix != "" {
		scope = fmt.Sprintf("route %q", rh.prefix)
	}

	for _, m := range mid {
		log.Printf("Applied middleware %q to %s", m.Name, scope)
	}
}

func (r *RouteHandler) Get(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}
func (r *RouteHandler) Post(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, pattern, handler)
}
func (r *RouteHandler) Put(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPut, pattern, handler)
}
func (r *RouteHandler) Delete(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}
