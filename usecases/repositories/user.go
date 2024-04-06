package repositories

import "github.com/ahmedshakshak/books-recommender/entities"

type User interface {
	IncreaseBookNumOfReadPages(*entities.Book) error
	SendMessage(*entities.User, string) error
}
