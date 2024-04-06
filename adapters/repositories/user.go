package repositories

import (
	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/models"
	"github.com/ahmedshakshak/books-recommender/notifiers"
)

// User implements the User interface in repositories by using a BookModelManager.
type User struct {
	bm models.BookModelManager
	n  notifiers.INotifier
}

// NewUser creates and returns a new instance of User, initializing it with the provided BookModelManager.
func NewUser(bm models.BookModelManager, notifier notifiers.INotifier) User {
	return User{
		bm: bm,
		n:  notifier,
	}
}

// IncreaseBookNumOfReadPages increase number of read pages for the provided book.
func (u *User) IncreaseBookNumOfReadPages(b *entities.Book) error {
	return u.bm.UpdateOrCreateBook(b)
}

// IncreaseBookNumOfReadPages increase number of read pages for the provided book.
func (u *User) SendMessage(user *entities.User, message string) error {
	return u.n.SendMessage(user.MobileNumber, message)
}
