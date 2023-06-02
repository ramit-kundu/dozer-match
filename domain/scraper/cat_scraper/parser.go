package catscraper

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseToChunk(ctx context.Context, html string) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	cardWrappers := make([]string, 0)

	doc.Find("div.card-wrapper").Each(func(i int, s *goquery.Selection) {
		outerHTML, err := s.Html()
		if err != nil {
			panic(err)
		}
		cardWrappers = append(cardWrappers, outerHTML)
	})

	err = ioutil.WriteFile("output.txt", []byte(cardWrappers[0]), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	GptParser(ctx, cardWrappers[0])
	fmt.Println("String successfully stored in file.")
}
