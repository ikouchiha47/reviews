package routes

import (
	"net/http"
	"reviews/http/httputils"
	"reviews/http/router"
)

type pingRouter struct {
	routes []router.Route
}

func NewPingRouter() *pingRouter {
	pingroutes := []router.Route{
		router.NewRouter(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
			return httputils.WriteJSON(w, 200, "pong")
		}),
	}

	return &pingRouter{routes: pingroutes}
}

func (pr *pingRouter) Routes() []router.Route {
	return pr.routes
}
