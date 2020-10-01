package dbstore

import (
	"context"
	"errors"
	"fmt"
	"reviews/errdefs"
	"reviews/internal/app/reviewsapp/entities"

	"github.com/jmoiron/sqlx"
)

type movieStore struct {
	ctx context.Context
	db  *sqlx.DB
}

var movies = []*entities.Movie{
	&entities.Movie{ID: "1", Name: "X-Men1", Rating: 2.4},
	&entities.Movie{ID: "2", Name: "X-Men2", Rating: 2.4},
	&entities.Movie{ID: "3", Name: "X-Men3", Rating: 2.4},
	&entities.Movie{ID: "4", Name: "X-Men4", Rating: 2.4},
}

func NewMovieStore(ctx context.Context, db *sqlx.DB) *movieStore {
	return &movieStore{ctx: ctx, db: db}
}

func (ms *movieStore) All() ([]*entities.Movie, error) {
	// s.db.ExecContext(...)
	return movies, nil
}

func (ms *movieStore) Find(ID string) (*entities.Movie, error) {
	for _, m := range movies {
		if m.ID == ID {
			return m, nil
		}
	}

	return nil, errdefs.NotFoundError(errors.New("Notfound"))
}

func (ms *movieStore) Create(movie entities.Movie) (*entities.Movie, error) {
	id := fmt.Sprintf("%d", len(movies)+1)

	newMovie := &entities.Movie{ID: id, Name: movie.Name, Rating: movie.Rating}
	movies = append(movies, newMovie)

	return newMovie, nil
}
