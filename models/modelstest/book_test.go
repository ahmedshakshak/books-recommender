package modelstest

import (
	"testing"

	"github.com/ahmedshakshak/books-recommender/entities"
	"github.com/ahmedshakshak/books-recommender/models"
	"github.com/ahmedshakshak/books-recommender/testhelpers"
)

func TestUpdateOrCreateBook(t *testing.T) {
	db := testhelpers.InitDB()
	defer testhelpers.TearDown(db)

	// Seed test data
	initialBook := &models.Book{Name: "Initial Book", NumOfReadPages: 0}
	db.Create(initialBook)

	// Run the test
	manager := models.NewBookManager(db)
	err := manager.UpdateOrCreateBook(&entities.Book{Id: initialBook.Id, NumOfReadPages: 10})
	if err != nil {
		t.Fatalf("failed to update or create book: %v", err)
	}

	// Verify the result
	var updatedBook models.Book
	db.First(&updatedBook, initialBook.Id)
	if updatedBook.NumOfReadPages != 10 {
		t.Errorf("expected NumOfReadPages to be 10, got %d", updatedBook.NumOfReadPages)
	}
}

func TestGetTopBooks(t *testing.T) {
	db := testhelpers.InitDB()
	defer testhelpers.TearDown(db)

	// Seed test data
	db.Create(&models.Book{Name: "Book 1", NumOfReadPages: 10})
	db.Create(&models.Book{Name: "Book 2", NumOfReadPages: 20})
	db.Create(&models.Book{Name: "Book 3", NumOfReadPages: 30})

	// Run the test
	manager := models.NewBookManager(db)
	books, err := manager.GetTopBooks("num_of_read_pages", 2)
	if err != nil {
		t.Fatalf("failed to get top books: %v", err)
	}

	// Verify the result
	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}

	if books[0].NumOfReadPages != int64(30) {
		t.Errorf("expected first book NumOfReadPages 30, got %d", books[0].NumOfReadPages)
	}

	if books[1].NumOfReadPages != int64(20) {
		t.Errorf("expected second book NumOfReadPages 20, got %d", books[1].NumOfReadPages)
	}
}
