package catscraper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	scraper "github.com/kundu-ramit/dozer_match/domain/dozer_crawler"
	"github.com/kundu-ramit/dozer_match/domain/entity"
)

type catScraper struct {
	parser Modularizer
}

func NewCatScraper() scraper.Scraper {
	return catScraper{
		parser: NewModularizer(false),
	}
}

func NewCatScraperGPT() scraper.Scraper {
	return catScraper{
		parser: NewModularizer(true),
	}
}

func (c catScraper) ScrapePage(ctx context.Context) ([]entity.BullDozer, error) {

	const baseUrl = "https://www.cat.com/en_US/products/new/equipment/dozers.html?page="

	// Launch headless browser
	launcher := launcher.New()
	url, err := launcher.Launch()
	if err != nil {
		return nil, err
	}

	dozers := []*entity.BullDozer{}

	// Connect to the browser and create a new page
	browser := rod.New().ControlURL(url)
	page := browser.MustConnect().MustPage()
	defer browser.MustClose()

	for i := 0; i < 2; i++ {
		res, err := c.scrapeSinglePage(ctx, page, i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dozers = append(dozers, res...)

	}
	return c.removeDuplicateAndFormat(dozers), nil
}

func (c catScraper) scrapeSinglePage(ctx context.Context, page *rod.Page, i int) ([]*entity.BullDozer, error) {
	// Navigate to a URL
	page.MustNavigate("https://www.cat.com/en_US/products/new/equipment/dozers.html?page=" + strconv.Itoa(i))

	waitDuration := time.Minute
	time.Sleep(waitDuration)

	// Extract page HTML
	html, err := page.HTML()
	if err != nil {
		return nil, err
	}
	return c.parser.Parse(ctx, html)
}

func (c catScraper) removeDuplicateAndFormat(dozers []*entity.BullDozer) []entity.BullDozer {
	hs := map[string]bool{}
	newDozer := []entity.BullDozer{}
	for i := 0; i < len(dozers); i++ {
		dozerHash := c.GenerateDozerHash(*dozers[i])
		_, exists := hs[dozerHash]
		if !exists {
			newDozer = append(newDozer, *dozers[i])
		}

	}
	return newDozer
}

func (c catScraper) GenerateDozerHash(dozer entity.BullDozer) string {
	return dozer.Make + " | " + dozer.Model + " | " + dozer.Picture + " | " + dozer.Category + " | " + dozer.EngineHP + " | " + strconv.Itoa(dozer.OperatingWeight)
}
