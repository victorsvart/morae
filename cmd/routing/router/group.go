package router

import (
	"log"
)

// Group represents a collection of routes sharing a common prefix and middleware
type Group struct {
	RouteHandler
}

// SubGroup creates a new nested route group with its own prefix and optional middleware.
func (g *Group) SubGroup(subGroupPrefix string, mid ...*Middleware) *SubGroup {
	subGroup := &SubGroup{
		RouteHandler: RouteHandler{router: g.router, prefix: g.prefix + subGroupPrefix},
		prefix:       subGroupPrefix,
	}

	log.Printf("Registered sub group: %s", subGroupPrefix)
	subGroup.Use(mid...)

	return subGroup
}
