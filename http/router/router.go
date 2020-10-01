package router

import "reviews/http/httputils"

type Router interface {
	Routes() []Route
}

type Route interface {
	Path() string
	Method() string
	Handler() httputils.ApiHandlerFunc
}

type localRouter struct {
	path, method string
	handler      httputils.ApiHandlerFunc
}

func (r localRouter) Path() string {
	return r.path
}

func (r localRouter) Method() string {
	return r.method
}

func (r localRouter) Handler() httputils.ApiHandlerFunc {
	return r.handler
}

func NewRouter(method, path string, handlerF httputils.ApiHandlerFunc) Route {
	var r = localRouter{method: method, path: path, handler: handlerF}

	return r
}
