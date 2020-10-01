package entities

import "time"

type Movie struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Rating    float64   `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
