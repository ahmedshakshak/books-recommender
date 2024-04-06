package repositories

import (
	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/models"
)

const (
	MOST_READ_BOOK_COLUMN_NAME = "num_of_read_pages"
)

// Book implements the Book interface in repositories by using a BookModelManager.
type Book struct {
	bm models.BookModelManager
}

// NewBook creates and returns a new instance of Book, initializing it with the provided BookModelManager.
func NewBook(bm models.BookModelManager) Book {
	return Book{
		bm: bm,
	}
}

// GetTheMostReadBooks top books based on number of read pages.
func (ba *Book) GetTheMostReadBooks(limit int64) ([]*entities.Book, error) {
	return ba.bm.GetTopBooks(MOST_READ_BOOK_COLUMN_NAME, int(limit))
}
