package dbstore

import (
	"context"
	"errors"
	"fmt"
	"reviews/errdefs"
	"reviews/internal/app/reviewsapp/entities"

	"github.com/jmoiron/sqlx"
)

type reviewsStore struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewReviewStore(ctx context.Context, db *sqlx.DB) *reviewsStore {
	return &reviewsStore{
		ctx: ctx,
		db:  db,
	}
}

var reviews = []*entities.Review{
	&entities.Review{ID: "1", MovieID: "1", Comment: "Waa", Like: true},
	&entities.Review{ID: "2", MovieID: "1", Comment: "Waa", Like: true},
	&entities.Review{ID: "3", MovieID: "2", Comment: "Waat", Like: true},
	&entities.Review{ID: "4", MovieID: "2", Comment: "Waat", Like: false},
	&entities.Review{ID: "5", MovieID: "3", Comment: "Waat the", Like: false},
	&entities.Review{ID: "6", MovieID: "3", Comment: "Waat the", Like: false},
	&entities.Review{ID: "7", MovieID: "4", Comment: "Damn", Like: true},
}

func (rs *reviewsStore) All() ([]*entities.Review, error) {
	// rs.db.ExecuteContext(s.ctx.....)

	return reviews, nil
}

func (rs *reviewsStore) FindByMovie(ID string) ([]*entities.Review, error) {
	reviews := []*entities.Review{}

	for _, review := range reviews {
		if review.MovieID == ID {
			reviews = append(reviews, review)
		}
	}

	return reviews, nil
}

func (rs *reviewsStore) FindBy(ID string) (*entities.Review, error) {
	for _, review := range reviews {
		if review.ID == ID {
			return review, nil
		}
	}

	return nil, errdefs.NotFoundError(errors.New("db error"))
}

func (rs *reviewsStore) Create(review entities.Review) (*entities.Review, error) {
	id := fmt.Sprintf("%d", len(reviews)+1)

	if review.MovieID == "" {
		return nil, errdefs.InvalidDataError()
	}

	newReview := &entities.Review{ID: id, MovieID: review.MovieID, Comment: review.Comment, Like: review.Like}

	reviews = append(reviews, newReview)

	return newReview, nil
}
