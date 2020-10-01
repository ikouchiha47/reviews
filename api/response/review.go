package response

import "time"

type ReviewResponse struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movie_id"`
	Comment   string    `json:"comment"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MovieReviewResponse struct {
	Movie   *MovieResponse    `json:"movie"`
	Reviews []*ReviewResponse `json:"reviews"`
}
