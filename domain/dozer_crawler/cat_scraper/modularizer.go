package catscraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kundu-ramit/dozer_match/domain/entity"
	"github.com/kundu-ramit/dozer_match/domain/scraper"
)

type Modularizer interface {
	Parse(ctx context.Context, html string) ([]*entity.BullDozer, error)
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

func (c modularizer) Parse(ctx context.Context, html string) ([]*entity.BullDozer, error) {

	dozers := []*entity.BullDozer{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	cardWrappers := make([]string, 0)

	doc.Find("div.card-wrapper").Each(func(i int, s *goquery.Selection) {
		outerHTML, err := s.Html()
		if err != nil {
			return
		}
		cardWrappers = append(cardWrappers, outerHTML)
	})

	for i := 0; i < len(cardWrappers); i++ {
		dozer, err := c.parser.Parse(ctx, cardWrappers[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		dozers = append(dozers, dozer)
	}

	return dozers, nil

}
