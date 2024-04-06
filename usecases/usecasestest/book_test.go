package usecasestest

import (
	"errors"
	"testing"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/mocks"
	"github.com/ahmedshakshak/books-recommender/usecases"
)

func TestRecommendBooks_Success(t *testing.T) {
	mockRepository := new(mocks.Book)
	expectedBooks := []*entities.Book{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}
	mockRepository.On("GetTheMostReadBooks", usecases.MAX_NUMBER_OF_RECOMMENDED_BOOKS).Return(expectedBooks, nil)
	b := usecases.NewBook(mockRepository)

	books, err := b.RecommendBooks()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(books) != len(expectedBooks) {
		t.Errorf("Expected %d books, got %d", len(expectedBooks), len(books))
	}

	mockRepository.AssertCalled(t, "GetTheMostReadBooks", usecases.MAX_NUMBER_OF_RECOMMENDED_BOOKS)
}

func TestRecommendBooks_Error(t *testing.T) {
	mockRepository := new(mocks.Book)
	mockRepository.On("GetTheMostReadBooks", usecases.MAX_NUMBER_OF_RECOMMENDED_BOOKS).Return(nil, errors.New("error getting most read books"))
	b := usecases.NewBook(mockRepository)

	_, err := b.RecommendBooks()

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	mockRepository.AssertCalled(t, "GetTheMostReadBooks", usecases.MAX_NUMBER_OF_RECOMMENDED_BOOKS)
}
