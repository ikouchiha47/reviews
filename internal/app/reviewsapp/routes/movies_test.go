package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reviews/api/response"
	"reviews/internal/app/reviewsapp/movies"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMovieRoutes_getMovies(t *testing.T) {
	t.Run("return 200 on success", func(t *testing.T) {

		movieService := new(movies.MockMovieService)
		movieService.On("AllMovies").Return([]*response.MovieResponse{}, nil)

		movieRouter := NewMovieRouter(movieService)
		movieRouter.InitRoutes()

		req, err := http.NewRequest("GET", "/movies", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		err = movieRouter.getMovies(rr, req, map[string]string{})

		require.NoError(t, err, "should not have failed to get result")

		movieService.AssertExpectations(t)
	})

	t.Run("return 500 on success", func(t *testing.T) {

		movieService := new(movies.MockMovieService)
		movieService.On("AllMovies").Return([]*response.MovieResponse{}, errors.New("bang"))

		movieRouter := NewMovieRouter(movieService)
		movieRouter.InitRoutes()

		req, err := http.NewRequest("GET", "/movies", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		err = movieRouter.getMovies(rr, req, map[string]string{})
		require.Error(t, err, "should have failed")

		movieService.AssertExpectations(t)
	})

}
