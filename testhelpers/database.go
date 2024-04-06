package testhelpers

import (
	"github.com/ahmedshakshak/books-recommender/models"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Connect to the test database
	db := models.GetDBConnection()
	err := db.AutoMigrate(models.Book{})
	if err != nil {
		panic(err)
	}

	return db
}

func TearDown(db *gorm.DB) {
	// Get all table names
	var tables []string
	result := db.Raw("SHOW TABLES").Scan(&tables)
	if result.Error != nil {
		panic(result.Error)
	}

	// Truncate each table
	for _, table := range tables {
		result := db.Exec("TRUNCATE TABLE " + table)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
