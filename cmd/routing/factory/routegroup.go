// Package factory provides route group registration logic used to organize application routes.
package factory

import "morae/cmd/routing/router"

// RouteGroup represents a modular group of routes that can be registered to a parent router group.
type RouteGroup interface {
	Register(*router.Group)
}
