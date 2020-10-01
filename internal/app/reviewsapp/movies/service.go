package movies

import (
	"reviews/api/response"
	"reviews/internal/app/reviewsapp/entities"
	"reviews/internal/app/reviewsapp/storage"
)

type MovieServiceInterface interface {
	AllMovies() ([]*response.MovieResponse, error)
	CreateMovie(movie entities.Movie) (*response.MovieResponse, error)
}

type movieService struct {
	store storage.MovieStore
}

func NewMovieService(mstore storage.MovieStore) *movieService {
	return &movieService{store: mstore}
}

func (ms *movieService) AllMovies() ([]*response.MovieResponse, error) {
	moviesResp := []*response.MovieResponse{}

	movies, err := ms.store.All()
	if err != nil {
		return moviesResp, err
	}

	for _, mov := range movies {
		moviesResp = append(moviesResp, MovieResponseFrom(mov))
	}

	return moviesResp, nil
}

func (ms *movieService) CreateMovie(movie entities.Movie) (*response.MovieResponse, error) {
	mov, err := ms.store.Create(entities.Movie{Name: movie.Name, Rating: 0})
	if err != nil {
		return nil, err
	}

	return MovieResponseFrom(mov), nil
}

func MovieResponseFrom(movie *entities.Movie) *response.MovieResponse {
	return &response.MovieResponse{
		ID:        movie.ID,
		Name:      movie.Name,
		Rating:    movie.Rating,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
	}
}
