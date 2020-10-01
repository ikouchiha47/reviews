package storage

import "reviews/internal/app/reviewsapp/entities"

type MovieStore interface {
	All() ([]*entities.Movie, error)
	Find(ID string) (*entities.Movie, error)
	Create(movie entities.Movie) (*entities.Movie, error)
}

type ReviewStore interface {
	All() ([]*entities.Review, error)
	FindBy(ID string) (*entities.Review, error)
	FindByMovie(ID string) ([]*entities.Review, error)
	Create(review entities.Review) (*entities.Review, error)
}
