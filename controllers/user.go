package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/usecases"
)

type User struct {
	uc usecases.IUser
}

func NewUser(uc usecases.IUser) User {
	return User{uc: uc}
}

func (u *User) SubmitReadingInterval(w http.ResponseWriter, r *http.Request) {

	body, err := u.parseSubmitReadingIntervalParams(r)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = u.uc.SubmitReadingInterval(
		&entities.User{Id: int64(body["user_id"].(float64))},
		&entities.Book{
			Id:             int64(body["book_id"].(float64)),
			NumOfReadPages: int64(body["end_page"].(float64)) - int64(body["start_page"].(float64)) + 1,
		},
	)

	if err != nil {
		log.Print(err)
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}
}

func (u *User) parseSubmitReadingIntervalParams(r *http.Request) (map[string]any, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("Missing request payload")
	}

	body := map[string]any{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// we can do some validation against the arguments as well and make sure all arguments exist

	return body, nil
}
