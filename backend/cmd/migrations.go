package cmd

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func ApplyMigration(db *gorm.DB) {

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{{
		// create `dozers` table
		ID: "201608301400_bull_dozers",
		Migrate: func(tx *gorm.DB) error {

			fmt.Println("Migrating 201608301400 create table")

			type bullDozer struct {
				Make            string `gorm:"column:make;"`
				Model           string `gorm:"column:model;"`
				Picture         string `gorm:"column:picture;"`
				Category        string `gorm:"column:category;"` //better as enum
				EngineHP        string `gorm:"column:engine_hp;"`
				OperatingWeight string `gorm:"column:operating_weight;"`
				ScrapeIndex     string `gorm:"column:scrape_index;"`
			}
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			return tx.Migrator().CreateTable(&bullDozer{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("bull_dozers")
		},
	}})

	if err := m.Migrate(); err != nil {
		panic(err)
	}
	log.Println("Migration did run successfully")
}
