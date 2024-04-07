package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/usecases"
)

type Book struct {
	bc usecases.IBook
}

func NewBook(bc usecases.IBook) Book {
	return Book{bc: bc}
}

func (b *Book) RecommendBooks(w http.ResponseWriter, r *http.Request) {
	books, err := b.bc.RecommendBooks()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}

	b.DecorateRecommendBooksResponse(w, books)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}
}

func (b *Book) DecorateRecommendBooksResponse(w http.ResponseWriter, books []*entities.Book) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(books)
}
