// Package router provides routing utilities including sub-group support for route handlers.
package router

// SubGroup represents a nested route handler group with a common prefix.
type SubGroup struct {
	RouteHandler        // Embedded route handler to inherit methods.
	prefix       string // Prefix used for all routes within this sub-group.
}
