package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Migrate performs database migration for the Book model.
func Migrate() {
	// Get a connection to the database
	db := GetDBConnection()

	// Migrate the schema to create the Book table in the database
	db.AutoMigrate(&Book{})
}

// GetDBConnection establishes a connection to the database.
// It returns a pointer to the gorm.DB object representing the database connection.
func GetDBConnection() *gorm.DB {
	MYSQL_HOSTNAME := os.Getenv("MYSQL_HOSTNAME")
	MYSQL_DB_NAME := os.Getenv("MYSQL_DB_NAME")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MYSQL_USERNAME,
		MYSQL_PASSWORD,
		MYSQL_HOSTNAME,
		MYSQL_PORT,
		MYSQL_DB_NAME,
	)

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Panic if there's an error connecting to the database
		panic("failed to connect database")
	}

	// Return the database connection
	return db
}
