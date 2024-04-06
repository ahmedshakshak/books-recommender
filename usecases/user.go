package usecases

import (
	"log"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/usecases/repositories"
)

const (
	SUBMISSION_MESSAGE = "Thanks for the submission"
)

type IUser interface {
	SubmitReadingInterval(user *entities.User, book *entities.Book) error
}

type user struct {
	repo repositories.User
}

func NewUser(repo repositories.User) IUser {
	return &user{repo: repo}
}

func (u *user) SubmitReadingInterval(user *entities.User, book *entities.Book) error {
	err := u.repo.IncreaseBookNumOfReadPages(book)
	if err != nil {
		return err
	}

	err = u.repo.SendMessage(user, SUBMISSION_MESSAGE)
	if err != nil {
		// handel the error when we fail to send a message since we already increase number of read pages for this book
		log.Print(err)
	}
	return nil
}
