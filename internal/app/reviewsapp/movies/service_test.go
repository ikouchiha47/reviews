package movies

import (
	"fmt"
	"reviews/errdefs"
	"reviews/internal/app/reviewsapp/entities"
	"reviews/internal/app/reviewsapp/storage/dbstore"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllMovies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		movieStore := new(dbstore.MockMovieStore)
		movieStore.On("All").Return([]*entities.Movie{}, nil)

		movieService := NewMovieService(movieStore)

		_, err := movieService.AllMovies()
		require.NoError(t, err, "should not have failed to get error")

		movieStore.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		movieStore := new(dbstore.MockMovieStore)
		movieStore.On("All").Return([]*entities.Movie{}, errdefs.InvalidDataError())

		movieService := NewMovieService(movieStore)

		_, err := movieService.AllMovies()
		require.Error(t, err, "should have failed to get data from store")

		fmt.Println(err)
		movieStore.AssertExpectations(t)
	})
}
