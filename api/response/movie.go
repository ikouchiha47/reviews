package response

import "time"

type MovieResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
