package main

import (
	"net/http"
	"strings"
)

type Group struct {
	router     *Router
	prefix     string
	middleware []Middleware
}

func (g *Group) Use(middleware ...Middleware) {
	g.middleware = append(g.middleware, middleware...)
}

func (g *Group) Handle(method, pattern string, handler http.HandlerFunc) {
	fullPattern := g.prefix
	if !strings.HasSuffix(g.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern
	g.router.Handle(method, fullPattern, handler)
}

func (g *Group) Get(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodGet, pattern, handler)
}

func (g *Group) Post(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodPost, pattern, handler)
}

func (g *Group) Put(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodPut, pattern, handler)
}

func (g *Group) Delete(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodDelete, pattern, handler)
}
