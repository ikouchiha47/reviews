package entities

import "time"

type Review struct {
	ID        string    `db:"id"`
	MovieID   string    `db:"movie_id"`
	Comment   string    `db:"name"`
	Like      bool      `db:"like"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
