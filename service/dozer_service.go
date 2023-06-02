package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/kundu-ramit/dozer_match/domain/cache"
	"github.com/kundu-ramit/dozer_match/domain/entity"
	"github.com/kundu-ramit/dozer_match/domain/repository"
	catscraper "github.com/kundu-ramit/dozer_match/domain/scraper/cat_scraper"
)

type DozerService interface {
	FetchDozers(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error)
}

type dozerService struct {
	repo    repository.BullDozerRepository
	cache   cache.Cache
	scraper catscraper.CatScraper
}

func NewDozerService() DozerService {
	return dozerService{
		repo:  repository.NewBullDozerRepository(),
		cache: cache.NewCache(),
	}
}

func (d dozerService) FetchDozers(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error) {
	res, _ := d.cache.Get(scrapeIndex)

	if res == "in_progress" {
		return nil, errors.New("scrape is in progress")
	}

	dozers, err := d.repo.Fetch(ctx, scrapeIndex)
	if err != nil {
		return nil, err
	}
	return dozers, nil
}

func (d dozerService) StartScrape(ctx context.Context) error {
	scrapeIndex := d.generateScrapeIndex()

	res, _ := d.cache.Get("https://www.cat.com")
	if res == "in_progress" {
		return errors.New("scrape is in progress")
	}

	d.cache.Set("https://www.cat.com", "in_progress", time.Hour)
	d.cache.Set(scrapeIndex, "in_progress", time.Hour)

	d.cache.Remove("https://www.cat.com")
	d.cache.Remove(scrapeIndex)

}

func (d dozerService) generateScrapeIndex() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	scrapeIndex := rand.Intn(max-min+1) + min

	return fmt.Sprintf("%010d", scrapeIndex)
}
