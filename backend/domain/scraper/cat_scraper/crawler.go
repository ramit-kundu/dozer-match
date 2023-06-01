package catscraper

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Crawl() {

	// Launch headless browser
	launcher := launcher.New()
	url, err := launcher.Launch()
	if err != nil {
		panic(err)
	}

	// Connect to the browser and create a new page
	browser := rod.New().ControlURL(url)
	page := browser.MustConnect().MustPage()
	defer browser.MustClose()

	// Navigate to a URL
	page.MustNavigate("https://www.cat.com/en_US/products/new/equipment/dozers.html?page=2")

	waitDuration := time.Minute
	time.Sleep(waitDuration)

	// Extract page HTML
	html, err := page.HTML()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", []byte(html), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("String successfully stored in file.")
}
