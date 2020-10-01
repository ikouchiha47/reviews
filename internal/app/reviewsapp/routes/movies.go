package routes

import (
	"net/http"
	"reviews/errdefs"
	"reviews/http/httputils"
	"reviews/http/router"
	"reviews/internal/app/reviewsapp/movies"
)

type movieRouter struct {
	service movies.MovieServiceInterface
	routes  []router.Route
}

func NewMovieRouter(svc movies.MovieServiceInterface) *movieRouter {
	mr := &movieRouter{service: svc}
	mr.InitRoutes()

	return mr
}

func (mr *movieRouter) Routes() []router.Route {
	return mr.routes
}

func (mr *movieRouter) InitRoutes() {
	mr.routes = []router.Route{
		router.NewRouter(http.MethodGet, "/movies", mr.getMovies),
	}
}

func (mr *movieRouter) getMovies(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	resp, err := mr.service.AllMovies()
	if err != nil {
		return errdefs.NewStatusError(500, err)
	}

	return httputils.WriteJSON(w, 200, resp)
}
