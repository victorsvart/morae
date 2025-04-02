package main

import (
	"log"
	"net/http"
	"strings"
)

type SubGroup struct {
	group      *Group
	prefix     string
	middleware []Middleware
}

func (g *Group) SubGroup(subGroupPrefix string) *SubGroup {
	subGroup := &SubGroup{
		group:  g,
		prefix: subGroupPrefix,
	}

	log.Printf("Registered sub group: %s", subGroupPrefix)
	return subGroup
}

func (g *SubGroup) Handle(method, pattern string, handler http.HandlerFunc) {
	fullPattern := g.prefix
	if !strings.HasSuffix(g.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern
	g.group.Handle(method, fullPattern, handler)
}

func (g *SubGroup) Get(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodGet, pattern, handler)
}

func (g *SubGroup) Post(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodPost, pattern, handler)
}

func (g *SubGroup) Put(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodPut, pattern, handler)
}

func (g *SubGroup) Delete(pattern string, handler http.HandlerFunc) {
	g.Handle(http.MethodDelete, pattern, handler)
}
