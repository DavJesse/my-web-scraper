package main

import (
	"context"
	"log"
	"strings"
	"time"

	"my-web-scraper/cmd/store"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

// func main() {
// 	response, err := http.Get("https://jiji.co.ke/")

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	var records []string

// 	records = append(records, "Title")

// 	doc, err := goquery.NewDocumentFromReader(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	doc.Find("div").Each(func(i int, s *goquery.Selection) {
// 		title := s.Text()
// 		// fmt.Printf("Title %d: %s\n", i+1, title)
// 		records = append(records, title)
// 	})
// 	store.SaveToJSON(records)
// }

func main() {
	// Create a new context for chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create new timeline to prevent script from hanging too long
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string // To store fully-rendered html

	// Navigate to URL to capture renderd html
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://jiji.co.ke/cars"),
		chromedp.WaitReady("body"),               // Wait 'til body is fully rendered
		chromedp.OuterHTML("html", &htmlContent), // Get full page in html
	)
	if err != nil {
		log.Fatal(err)
	}

	var records []string

	// Parse fully renderd html
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		// fmt.Printf("Title %d: %s\n", i+1, title)
		records = append(records, title)
	})
	store.SaveToJSON(records)
}
