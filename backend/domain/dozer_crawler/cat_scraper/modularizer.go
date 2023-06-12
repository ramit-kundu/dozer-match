package catscraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	scraper "github.com/kundu-ramit/dozer_match/domain/dozer_crawler"
	"github.com/kundu-ramit/dozer_match/domain/entity"
)

type Modularizer interface {
	ParseModule(ctx context.Context, html string) ([]*entity.BullDozer, error)
}
type modularizer struct {
	parser scraper.Parser
}

func NewModularizer(useGpt bool) Modularizer {
	if useGpt {
		return modularizer{
			parser: NewGptParser(),
		}
	} else {
		return modularizer{
			parser: NewManualParser(),
		}
	}
}

func (c modularizer) ParseModule(ctx context.Context, html string) ([]*entity.BullDozer, error) {

	dozers := []*entity.BullDozer{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error happened while extracting document" + err.Error())
		return nil, err
	}

	cardWrappers := make([]string, 0)

	doc.Find("div.card-wrapper").Each(func(i int, s *goquery.Selection) {
		outerHTML, err := s.Html()
		if err != nil {
			fmt.Println("error happened while getting html" + err.Error())
			return
		}
		cardWrappers = append(cardWrappers, outerHTML)
	})

	for i := 0; i < len(cardWrappers); i++ {
		dozer, err := c.parser.Parse(ctx, cardWrappers[i])
		if err != nil {
			fmt.Println("error happened while parsing chunk of html" + err.Error())
			continue
		}
		dozers = append(dozers, dozer)
	}

	return dozers, nil

}
