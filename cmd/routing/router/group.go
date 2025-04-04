package router

import (
	"log"
)

type Group struct {
	RouteHandler
	middleware []Middleware
}

func (g *Group) SubGroup(subGroupPrefix string, mid ...*Middleware) *SubGroup {
	subGroup := &SubGroup{
		RouteHandler: RouteHandler{router: g.router, prefix: g.prefix + subGroupPrefix},
	}

	log.Printf("Registered sub group: %s", subGroupPrefix)
	subGroup.Use(mid...)

	return subGroup
}
