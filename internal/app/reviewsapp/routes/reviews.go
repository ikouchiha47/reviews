package routes

import (
	"net/http"
	"reviews/errdefs"
	"reviews/http/httputils"
	"reviews/http/router"
	"reviews/internal/app/reviewsapp/entities"
	"reviews/internal/app/reviewsapp/reviews"
)

type movieReviewRouter struct {
	service reviews.MovieReviewInterface
	routes  []router.Route
}

func NewMovieReviewRouter(svc reviews.MovieReviewInterface) *movieReviewRouter {
	mrr := &movieReviewRouter{service: svc}
	mrr.InitRoutes()

	return mrr
}

func (mrr *movieReviewRouter) createReview(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	rev, err := httputils.ReadJSON(r, entities.Review{})
	if err != nil {
		return err
	}

	review := rev.(entities.Review)

	resp, errr := mrr.service.CreateReview(review)
	if err != nil {
		return errdefs.NewStatusError(500, errr)
	}

	return httputils.WriteJSON(w, 201, resp)
}

func (mrr *movieReviewRouter) InitRoutes() {
	mrr.routes = []router.Route{
		router.NewRouter(http.MethodGet, "/movies/reviews", mrr.getMovieReview),
		router.NewRouter(http.MethodPost, "/reviews", mrr.createReview),
	}
}

func (mrr *movieReviewRouter) getMovieReview(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	resp, err := mrr.service.AllMovieReviews()
	if err != nil {
		return errdefs.NewStatusError(500, err)
	}

	return httputils.WriteJSON(w, 200, resp)
}

func (mrr *movieReviewRouter) Routes() []router.Route {
	return mrr.routes
}
