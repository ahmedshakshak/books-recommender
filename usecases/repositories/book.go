package repositories

import "github.com/ahmedshakshak/books-recommender/entities"

type Book interface {
	GetTheMostReadBooks(numOfBooks int64) ([]*entities.Book, error)
}
