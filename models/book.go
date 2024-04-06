package models

import (
	"fmt"

	"github.com/ahmedshakshak/books-recommender/entities"
	"gorm.io/gorm"
)

// Book represents the model for a book entity in the database.
type Book struct {
	Id             int64 `gorm:"primaryKey;autoIncrement"`
	Name           string
	NumOfReadPages int64
}

// BookModelManager manages database operations related to the Book model.
type BookModelManager struct {
	db *gorm.DB
}

// NewBookManager creates and returns a new instance of BookModelManager, initializing it with the provided gorm.DB.
func NewBookManager(db *gorm.DB) BookModelManager {
	return BookModelManager{
		db: db,
	}
}

// UpdateBook updates the information of an existing book in the database, if it doesn't exist it creates it.
func (bm *BookModelManager) UpdateOrCreateBook(b *entities.Book) error {
	model := Book{
		Id:             b.Id,
		Name:           b.Name,
		NumOfReadPages: b.NumOfReadPages,
	}
	// try to update it the row
	tx := bm.db.Model(&model).Where("id = ?", model.Id).UpdateColumn("num_of_read_pages", gorm.Expr("num_of_read_pages + ?", model.NumOfReadPages))
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 { // row doesn't exist
		tx = bm.db.Create(&model)
	}

	return tx.Error
}

// GetXopBooks get 'x' books from the table sorted by 'col' desc.
func (bm *BookModelManager) GetTopBooks(col string, limit int) ([]*entities.Book, error) {
	models := make([]Book, 0)

	bm.db.Order(fmt.Sprintf("%s desc", col)).Limit(limit).Find(&models)

	return bookModelsToEntities(models), nil
}

// bookModelToEntity converts a Book model to a Book entity.
func bookModelToEntity(model *Book) *entities.Book {
	return &entities.Book{
		Id:             model.Id,
		Name:           model.Name,
		NumOfReadPages: model.NumOfReadPages,
	}
}

// bookModelsToEntities converts a slice of Book models to a slice of Book entities.
func bookModelsToEntities(models []Book) []*entities.Book {
	books := make([]*entities.Book, len(models))

	for i := 0; i < len(books); i++ {
		books[i] = bookModelToEntity(&models[i])
	}

	return books
}
