package service

import (
	"context"
	"errors"
	"fmt"
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
	StartScrape(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error)
	CheckExistingScrape(ctx context.Context) error
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

func MakeNewDozerService(repo repository.BullDozerRepository, cache cache.Cache) DozerService {
	var sc scraper.Scraper
	if os.Getenv("USE_GPT") == "true" {
		sc = catscraper.NewCatScraperGPT()
	} else {
		sc = catscraper.NewCatScraper()
	}
	return dozerService{
		repo:    repo,
		cache:   cache,
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
	if err != nil || scrapeIndex == "" {
		return nil, errors.New("error in getting latest scrape index maybe no scraping is done")
	}

	dozers, err := d.repo.Fetch(ctx, scrapeIndex)
	if err != nil {
		return nil, err
	}
	return dozers, nil
}

func (d dozerService) CheckExistingScrape(ctx context.Context) error {
	res, _ := d.cache.Get("catdotcom")
	if res == "in_progress" {
		fmt.Println("scrape is in progress")
		return errors.New("scrape is in progress")
	}
	return nil
}

func (d dozerService) StartScrape(ctx context.Context, scrapeIndex string) ([]entity.BullDozer, error) {

	d.cache.Set("catdotcom", "in_progress", 3*time.Minute)
	d.cache.Set(scrapeIndex, "in_progress", 3*time.Minute)

	dozers, err := d.scraper.ScrapePage(ctx)

	if err != nil {
		d.cache.Remove("catdotcom")
		d.cache.Remove(scrapeIndex)
		return nil, err
	}

	for i := 0; i < len(dozers); i++ {
		dozers[i].ScrapeIndex = scrapeIndex
	}

	d.repo.BulkCreate(ctx, dozers)

	d.cache.Remove("catdotcom")
	d.cache.Remove(scrapeIndex)
	d.cache.Set("cat_latest_index", scrapeIndex, 0)

	return dozers, nil

}

func (d dozerService) Delete(ctx context.Context) error {
	d.cache.FlushAll()
	return d.repo.Delete(ctx)
}
