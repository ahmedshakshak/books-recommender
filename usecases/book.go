package usecases

import (
	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/usecases/repositories"
)

const (
	MAX_NUMBER_OF_RECOMMENDED_BOOKS int64 = 5
)

type IBook interface {
	RecommendBooks() ([]*entities.Book, error)
}

type book struct {
	repo repositories.Book
}

func NewBook(repo repositories.Book) IBook {
	return &book{repo: repo}
}

func (b *book) RecommendBooks() ([]*entities.Book, error) {
	return b.repo.GetTheMostReadBooks(MAX_NUMBER_OF_RECOMMENDED_BOOKS)
}
