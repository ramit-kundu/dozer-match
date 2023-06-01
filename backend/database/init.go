package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := "host=database-host user=user password=password dbname=database-name port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// You can add GORM configurations or customizations here

	return db, nil
}

func main() {
	db, err := InitializeDatabase()
	if err != nil {
		// Handle the error
	}

	// Use the db connection for your database operations
	// ...

	// Don't forget to close the connection when you're done
	db.Close()
}
