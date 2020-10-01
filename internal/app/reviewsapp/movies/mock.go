package movies

import (
	"reviews/api/response"
	"reviews/internal/app/reviewsapp/entities"

	"github.com/stretchr/testify/mock"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) AllMovies() ([]*response.MovieResponse, error) {
	args := m.Called()

	return args.Get(0).([]*response.MovieResponse), args.Error(1)
}

func (m *MockMovieService) CreateMovie(movie entities.Movie) (*response.MovieResponse, error) {
	args := m.Called()

	return args.Get(0).(*response.MovieResponse), args.Error(1)
}
