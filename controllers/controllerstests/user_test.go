package controllerstests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmedshakshak/books-recommender/controllers"
	"github.com/ahmedshakshak/books-recommender/mocks"
	"github.com/stretchr/testify/mock"
)

func TestSubmitReadingInterval(t *testing.T) {
	userUsecaseMock := &mocks.IUser{}
	user := controllers.NewUser(userUsecaseMock)

	// Mocking the behavior of SubmitReadingInterval method
	userUsecaseMock.On("SubmitReadingInterval", mock.Anything, mock.Anything).Return(nil)

	reqBody := map[string]interface{}{
		"user_id":    123,
		"book_id":    456,
		"start_page": 10,
		"end_page":   20,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/submitReadingInterval", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.SubmitReadingInterval)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert that the mock's method was called with expected arguments
	userUsecaseMock.AssertCalled(t, "SubmitReadingInterval", mock.Anything, mock.Anything)
}

func TestSubmitReadingInterval_InvalidPayload(t *testing.T) {
	userUsecaseMock := &mocks.IUser{}
	user := controllers.NewUser(userUsecaseMock)

	req, err := http.NewRequest("POST", "/submitReadingInterval", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.SubmitReadingInterval)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSubmitReadingInterval_UsecaseError(t *testing.T) {
	userUsecaseMock := &mocks.IUser{}
	user := controllers.NewUser(userUsecaseMock)

	// Mocking the behavior of SubmitReadingInterval method to return an error
	userUsecaseMock.On("SubmitReadingInterval", mock.Anything, mock.Anything).Return(errors.New("usecase error"))

	reqBody := map[string]interface{}{
		"user_id":    123,
		"book_id":    456,
		"start_page": 10,
		"end_page":   20,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/submitReadingInterval", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.SubmitReadingInterval)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Assert that the mock's method was called with expected arguments
	userUsecaseMock.AssertCalled(t, "SubmitReadingInterval", mock.Anything, mock.Anything)
}
