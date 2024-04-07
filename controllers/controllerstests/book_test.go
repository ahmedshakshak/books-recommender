package controllerstests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmedshakshak/books-recommender/controllers"
	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/mocks"
)

func TestRecommendBooks(t *testing.T) {
	bookUsecaseMock := &mocks.IBook{}
	book := controllers.NewBook(bookUsecaseMock)

	// Mocking the behavior of RecommendBooks method
	bookUsecaseMock.On("RecommendBooks").Return([]*entities.Book{}, nil)

	req, err := http.NewRequest("GET", "/recommendBooks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(book.RecommendBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Assert that the mock's method was called
	bookUsecaseMock.AssertExpectations(t)
}

func TestRecommendBooks_UsecaseError(t *testing.T) {
	bookUsecaseMock := &mocks.IBook{}
	book := controllers.NewBook(bookUsecaseMock)

	// Mocking the behavior of RecommendBooks method to return an error
	bookUsecaseMock.On("RecommendBooks").Return(nil, errors.New("usecase error"))

	req, err := http.NewRequest("GET", "/recommendBooks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(book.RecommendBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Assert that the mock's method was called
	bookUsecaseMock.AssertExpectations(t)
}
