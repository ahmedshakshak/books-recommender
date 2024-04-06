package usecasestest

import (
	"errors"
	"testing"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/mocks"
	"github.com/ahmedshakshak/books-recommender/usecases"
)

func TestSubmitReadingInterval_Success(t *testing.T) {
	mockRepository := new(mocks.User)
	book := &entities.Book{}
	mockRepository.On("IncreaseBookNumOfReadPages", book).Return(nil)
	user := &entities.User{}
	mockRepository.On("SendMessage", user, usecases.SUBMISSION_MESSAGE).Return(nil)
	u := usecases.NewUser(mockRepository)

	err := u.SubmitReadingInterval(user, book)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockRepository.AssertCalled(t, "IncreaseBookNumOfReadPages", book)
	mockRepository.AssertCalled(t, "SendMessage", user, usecases.SUBMISSION_MESSAGE)
}

func TestSubmitReadingInterval_IncreaseBookNumOfReadPagesError(t *testing.T) {
	mockRepository := new(mocks.User)
	book := &entities.Book{}
	mockRepository.On("IncreaseBookNumOfReadPages", book).Return(errors.New("error increasing read pages"))
	user := &entities.User{}
	u := usecases.NewUser(mockRepository)

	err := u.SubmitReadingInterval(user, book)

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	mockRepository.AssertCalled(t, "IncreaseBookNumOfReadPages", book)
	mockRepository.AssertNotCalled(t, "SendMessage", user, usecases.SUBMISSION_MESSAGE)
}

func TestSubmitReadingInterval_SendMessageError(t *testing.T) {
	mockRepository := new(mocks.User)
	book := &entities.Book{}
	mockRepository.On("IncreaseBookNumOfReadPages", book).Return(nil)
	user := &entities.User{}
	mockRepository.On("SendMessage", user, usecases.SUBMISSION_MESSAGE).Return(errors.New("error sending message"))
	u := usecases.NewUser(mockRepository)

	err := u.SubmitReadingInterval(user, book)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockRepository.AssertCalled(t, "IncreaseBookNumOfReadPages", book)
	mockRepository.AssertCalled(t, "SendMessage", user, usecases.SUBMISSION_MESSAGE)
}
