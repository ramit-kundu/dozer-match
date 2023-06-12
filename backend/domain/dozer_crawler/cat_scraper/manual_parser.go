package catscraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	scraper "github.com/kundu-ramit/dozer_match/domain/dozer_crawler"
	"github.com/kundu-ramit/dozer_match/domain/entity"
)

type manualParser struct{}

func NewManualParser() scraper.Parser {
	return manualParser{}
}

func (m manualParser) Parse(ctx context.Context, html string) (*entity.BullDozer, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error happened during reading document" + err.Error())
		return nil, err
	}

	Make := doc.Find(".card").AttrOr("data-brand", "")
	model := doc.Find(".value.spec:contains('Engine Model')").Next().Text()
	picture, _ := doc.Find("img.lazyloaded").Attr("src")
	category := doc.Find(".value.family").Text()
	engineHP := doc.Find(".value.spec:contains('Power')").Next().Text()
	operatingWeight := doc.Find(".value.spec:contains('Weight')").Next().Text()

	if strings.Contains(operatingWeight, "lb") {
		operatingWeight = strings.Split(operatingWeight, " ")[0]
	}
	if strings.Contains(engineHP, "HP") {
		engineHP = strings.Split(engineHP, " ")[0]
	}

	return &entity.BullDozer{
		Make:            Make,
		Model:           model,
		Picture:         picture,
		Category:        category,
		EngineHP:        engineHP,
		OperatingWeight: operatingWeight,
	}, nil
}
