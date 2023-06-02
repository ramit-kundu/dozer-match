package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() *gorm.DB {
	dsn := "user=root1 password=sample host=localhost port=5432 dbname=bb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
