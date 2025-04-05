package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RouteHandler struct {
	router     *Router
	middleware []*Middleware
	prefix     string
}

func (rh *RouteHandler) Handle(method, pattern string, handler http.HandlerFunc, middlewareExclude ...*string) {
	fullPattern := rh.prefix
	if !strings.HasSuffix(rh.prefix, "/") && !strings.HasPrefix(pattern, "/") && pattern != "" {
		fullPattern = "/"
	}

	fullPattern += pattern
	for _, mw := range rh.middleware {
		skip := false
		for _, excluded := range middlewareExclude {
			if *excluded == mw.Name {
				skip = true
				break
			}
		}

		if skip {
			continue
		}
		handler = mw.Exec(handler)
	}
	rh.router.Handle(method, fullPattern, handler)
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

func (r *RouteHandler) Get(pattern string, handler http.HandlerFunc, middlewareExclude ...*string) {
	r.Handle(http.MethodGet, pattern, handler, middlewareExclude...)
}

func (r *RouteHandler) Post(pattern string, handler http.HandlerFunc, middlewareExclude ...*string) {
	r.Handle(http.MethodPost, pattern, handler)
}

func (r *RouteHandler) Put(pattern string, handler http.HandlerFunc, middlewareExclude ...*string) {
	r.Handle(http.MethodPut, pattern, handler)
}

func (r *RouteHandler) Delete(pattern string, handler http.HandlerFunc, middlewareExclude ...*string) {
	r.Handle(http.MethodDelete, pattern, handler)
}
