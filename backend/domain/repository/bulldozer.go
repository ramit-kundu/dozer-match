package repository

import (
	"context"

	"github.com/kundu-ramit/dozer_match/domain/entity"
	"github.com/kundu-ramit/dozer_match/infra/database"
	"gorm.io/gorm"
)

type BullDozerRepository interface {
	BulkCreate(ctx context.Context, dozers []entity.BullDozer) error
	Fetch(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error)
	Delete(ctx context.Context) error
}

type bullDozerRepository struct {
	db *gorm.DB
}

func NewBullDozerRepository() BullDozerRepository {
	return &bullDozerRepository{
		db: database.Initialize(),
	}
}

func (r bullDozerRepository) BulkCreate(ctx context.Context, dozers []entity.BullDozer) error {
	err := r.db.Create(&dozers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r bullDozerRepository) Fetch(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error) {
	var dozers []entity.BullDozer
	err := r.db.Where("scrape_index = ?", scrapeIndex).Find(&dozers).Error
	if err != nil {
		return nil, err
	}
	return dozers, nil
}

func (r bullDozerRepository) Delete(ctx context.Context) error {
	//gorm needs a where clause
	err := r.db.Where("1 = 1").Delete(&entity.BullDozer{}, "").Error
	if err != nil {
		return err
	}
	return nil
}
