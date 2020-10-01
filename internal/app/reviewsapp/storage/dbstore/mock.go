// Package dbstore provides ...
package dbstore

import (
	"reviews/internal/app/reviewsapp/entities"

	"github.com/stretchr/testify/mock"
)

type MockMovieStore struct {
	mock.Mock
}

func (m *MockMovieStore) All() ([]*entities.Movie, error) {
	args := m.Called()

	return args.Get(0).([]*entities.Movie), args.Error(1)
}

func (m *MockMovieStore) Find(ID string) (*entities.Movie, error) {
	args := m.Called(ID)

	return args.Get(0).(*entities.Movie), args.Error(1)
}

func (m *MockMovieStore) Create(movie entities.Movie) (*entities.Movie, error) {
	args := m.Called(movie)

	return args.Get(0).(*entities.Movie), args.Error(1)
}

type MockReviewStore struct {
	mock.Mock
}

func (m *MockReviewStore) All() ([]*entities.Review, error) {
	args := m.Called()

	return args.Get(0).([]*entities.Review), args.Error(1)
}

func (m *MockReviewStore) FindBy(ID string) (*entities.Review, error) {
	args := m.Called(ID)

	return args.Get(0).(*entities.Review), args.Error(1)
}

func (m *MockReviewStore) FindByMovie(ID string) ([]*entities.Review, error) {
	args := m.Called(ID)

	return args.Get(0).([]*entities.Review), args.Error(1)
}

func (m *MockReviewStore) Create(review entities.Review) (*entities.Review, error) {
	args := m.Called(review)

	return args.Get(0).(*entities.Review), args.Error(1)
}
