package factory

import "morae/cmd/routing/router"

type RouteGroup interface {
  Register(*router.Group)
}
