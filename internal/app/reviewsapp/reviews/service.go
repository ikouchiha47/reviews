package reviews

import (
	"reviews/api/response"
	"reviews/internal/app/reviewsapp/entities"
	"reviews/internal/app/reviewsapp/storage"
)

type ReviewServiceInterface interface {
	FindByMovie(ID string) ([]*response.ReviewResponse, error)
	CreateReview(review entities.Review) (*response.ReviewResponse, error)
}

type reviewService struct {
	store storage.ReviewStore
}

func NewReviewService(rstore storage.ReviewStore) *reviewService {
	return &reviewService{store: rstore}
}

func (rs *reviewService) FindByMovie(ID string) ([]*response.ReviewResponse, error) {
	reviewsResp := []*response.ReviewResponse{}

	reviews, err := rs.store.FindByMovie(ID)
	if err != nil {
		return reviewsResp, err
	}

	for _, review := range reviews {
		reviewsResp = append(reviewsResp, ReviewResponseFrom(review))
	}

	return reviewsResp, nil
}

func (rs *reviewService) CreateReview(review entities.Review) (*response.ReviewResponse, error) {
	r, err := rs.store.Create(review)
	if err != nil {
		return nil, err
	}

	return ReviewResponseFrom(r), nil
}

func ReviewResponseFrom(review *entities.Review) *response.ReviewResponse {
	return &response.ReviewResponse{
		ID:        review.ID,
		MovieID:   review.MovieID,
		Comment:   review.Comment,
		Like:      review.Like,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}
