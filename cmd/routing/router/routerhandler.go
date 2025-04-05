package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"slices"
)

type RouteHandler struct {
	router     *Router
	middleware []*Middleware
	prefix     string
}

func (rh *RouteHandler) Handle(method, pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	fullPattern := rh.prefix
	if !strings.HasSuffix(rh.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern

  // if not excluded from middleware execute it otherwise skip
  for _, mw := range rh.middleware {
    if opts != nil && slices.Contains(opts.MiddlewareExclude, mw.Name) {
      log.Printf("Route %s %s will excluded from %s", method, fullPattern, mw.Name)
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

func (r *RouteHandler) Get(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	r.Handle(http.MethodGet, pattern, handler, opts)
}

func (r *RouteHandler) Post(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	r.Handle(http.MethodPost, pattern, handler, opts)
}

func (r *RouteHandler) Put(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	r.Handle(http.MethodPut, pattern, handler, opts)
}

func (r *RouteHandler) Delete(pattern string, handler http.HandlerFunc, opts *RouteOptions) {
	r.Handle(http.MethodDelete, pattern, handler, opts)
}
