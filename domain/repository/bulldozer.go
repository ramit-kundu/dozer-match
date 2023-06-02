package repository

import (
	"github.com/kundu-ramit/dozer_match/domain/entity"
	"gorm.io/gorm"
)

type BullDozerRepository interface {
	BulkCreate(dozers []entity.BullDozer) error
	Fetch(scrapeIndex int64) ([]entity.BullDozer, error)
}

type bullDozerRepository struct {
	db *gorm.DB
}

func NewBullDozerRepository(db *gorm.DB) BullDozerRepository {
	return &bullDozerRepository{
		db: db,
	}
}

func (r bullDozerRepository) BulkCreate(dozers []entity.BullDozer) error {
	err := r.db.Create(&dozers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r bullDozerRepository) Fetch(scrapeIndex int64) ([]entity.BullDozer, error) {
	var dozers []entity.BullDozer
	err := r.db.Where("scrape_index = ?", scrapeIndex).Find(&dozers).Error
	if err != nil {
		return nil, err
	}
	return dozers, nil
}
