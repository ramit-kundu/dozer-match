package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/kundu-ramit/dozer_match/domain/cache"
	scraper "github.com/kundu-ramit/dozer_match/domain/dozer_crawler"
	catscraper "github.com/kundu-ramit/dozer_match/domain/dozer_crawler/cat_scraper"
	"github.com/kundu-ramit/dozer_match/domain/entity"
	"github.com/kundu-ramit/dozer_match/domain/repository"
)

type DozerService interface {
	FetchById(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error)
	FetchLatest(ctx context.Context) ([]entity.BullDozer, error)
	StartScrape(ctx context.Context, scrapeIndex string) error
	Delete(ctx context.Context) error
}

type dozerService struct {
	repo    repository.BullDozerRepository
	cache   cache.Cache
	scraper scraper.Scraper
}

func NewDozerService() DozerService {
	var sc scraper.Scraper
	if os.Getenv("USE_GPT") == "true" {
		sc = catscraper.NewCatScraperGPT()
	} else {
		sc = catscraper.NewCatScraper()
	}
	return dozerService{
		repo:    repository.NewBullDozerRepository(),
		cache:   cache.NewCache(),
		scraper: sc,
	}
}

func (d dozerService) FetchById(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error) {
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

func (d dozerService) FetchLatest(ctx context.Context) ([]entity.BullDozer, error) {

	scrapeIndex, err := d.cache.Get("cat_latest_index")
	if err != nil {
		return nil, errors.New("error in getting latest scrape index maybe no scraping is done")
	}

	dozers, err := d.repo.Fetch(ctx, scrapeIndex)
	if err != nil {
		return nil, err
	}
	return dozers, nil
}

func (d dozerService) StartScrape(ctx context.Context, scrapeIndex string) error {

	res, _ := d.cache.Get("https://www.cat.com")
	if res == "in_progress" {
		return errors.New("scrape is in progress")
	}

	d.cache.Set("https://www.cat.com", "in_progress", time.Hour)
	d.cache.Set(scrapeIndex, "in_progress", time.Hour)

	dozers, err := d.scraper.ScrapePage(ctx)

	if err != nil {
		return err
	}
	d.repo.BulkCreate(ctx, dozers)

	d.cache.Remove("https://www.cat.com")
	d.cache.Remove(scrapeIndex)
	d.cache.Set("cat_latest_index", scrapeIndex, 0)

	return nil

}

func (d dozerService) Delete(ctx context.Context) error {
	return d.repo.Delete(ctx)
}
