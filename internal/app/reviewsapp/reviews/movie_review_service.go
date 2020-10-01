package reviews

import (
	"reviews/api/response"
	"reviews/internal/app/reviewsapp/entities"
	"reviews/internal/app/reviewsapp/movies"
)

type MovieReviewInterface interface {
	AllMovieReviews() ([]*response.MovieReviewResponse, error)
	CreateReview(review entities.Review) (*response.ReviewResponse, error)
}

type MovieReviewService struct {
	MovieService  movies.MovieServiceInterface
	ReviewService ReviewServiceInterface
}

func (mrs *MovieReviewService) AllMovieReviews() ([]*response.MovieReviewResponse, error) {
	movieReviews := []*response.MovieReviewResponse{}

	movies, err := mrs.MovieService.AllMovies()
	if err != nil {
		return movieReviews, err
	}

	for _, movie := range movies {
		reviews, err := mrs.ReviewService.FindByMovie(movie.ID)
		if err != nil {
			return movieReviews, err
		}

		movieReview := &response.MovieReviewResponse{
			Movie:   movie,
			Reviews: reviews,
		}

		movieReviews = append(movieReviews, movieReview)
	}

	return movieReviews, nil
}

func (mrs *MovieReviewService) CreateReview(review entities.Review) (*response.ReviewResponse, error) {
	return mrs.ReviewService.CreateReview(review)
}
