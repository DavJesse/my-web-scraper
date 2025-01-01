package services

import (
	"fmt"
	"my-web-scraper/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeDataWithGoQuery(path, targetDiv string) ([]models.CarListing, error) {
	var listings []models.CarListing

	// Make HTTP request to the provided path
	response, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("failed to complete http request: %v", err)
	}
	defer response.Body.Close()

	// Retrieve the document from the response body
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch document: %v", err)
	}

	// Extract car listings from the document by targeting specific class names
	doc.Find(targetDiv).Each(func(i int, s *goquery.Selection) {
		listing := models.CarListing{}

		// Extract title, Price, Description, and Location
		listing.Title = strings.TrimSpace(s.Find(".b-advert-title-inner.qa-advert-title.b-advert-title-inner--div").Text())
		listing.Price = strings.TrimSpace(s.Find(".qa-advert-price").Text())
		listing.Description = strings.TrimSpace(s.Find(".b-list-advert-base__description-text").Text())
		listing.Location = strings.TrimSpace(s.Find(".b-list-advert__region__text").Text())
		listing.Condition = strings.TrimSpace(s.Find(".b-list-advert-base__item-attr").Text())

		listings = append(listings, listing)
	})
}

func ScrapeDataWithHeadless(path, targetDiv string) ([]models.CarListing, error) {
	// Launch headless browser
	ctx, cancel := LaunchHeadlessBrowser(path, targetDiv)
	if ctx == nil {
		log.Fatal("Failed to launch browser")
	}
	defer cancel()

	var listings []models.CarListing

	// Navigate to web page
	err := chromedp.Run(ctx,
		chromedp.Navigate(path),
		chromedp.WaitVisible(targetDiv, chromedp.ByQuery),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to navigate to webpage: %v", err)
	}

	for len(listings) < 1000 {
		// Scroll to bottom of the page to load more
		err = chromedp.Run(ctx,
			chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight);`, nil),
			chromedp.Sleep(2*time.Second), // Wait for content to load
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scroll: %v", err)
		}

		// Find all car listings on the page
		var elements []*cdp.Node
		err = chromedp.Run(ctx,
			chromedp.Nodes(".masonry-item", &elements, chromedp.ByQueryAll),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to find elements: %v", err)
		}

		newListings := len(elements)
		if newListings <= len(listings) {
			break
		}

		// Scrape listings
		for i := len(listings); i < newListings; i++ {
			listing := models.CarListing{}

			err = chromedp.Run(ctx,
				chromedp.Text(".b-advert-title-inner.qa-advert-title.b-advert-title-inner--div", &listing.Title, chromedp.ByQuery, chromedp.FromNode(elements[i])),
				chromedp.Text(".qa-advert-price", &listing.Price, chromedp.ByQuery, chromedp.FromNode(elements[i])),
				chromedp.Text(".b-list-advert-base__description-text", &listing.Description, chromedp.ByQuery, chromedp.FromNode(elements[i])),
				chromedp.Text(".b-list-advert__region__text", &listing.Location, chromedp.ByQuery, chromedp.FromNode(elements[i])),
				chromedp.Text(".b-list-advert-base__item-attr", &listing.Condition, chromedp.ByQuery, chromedp.FromNode(elements[i])),
			)
			if err != nil {
				log.Printf("Error scraping listing: %v", err)
				continue
			}

			// Trim leading and trailing whitespaces from fields for uniformity/readability
			listing.Title = strings.TrimSpace(listing.Title)
			listing.Price = strings.TrimSpace(listing.Price)
			listing.Description = strings.TrimSpace(listing.Description)
			listing.Location = strings.TrimSpace(listing.Location)
			listing.Condition = strings.TrimSpace(listing.Condition)

			listings = append(listings, listing)
		}

		log.Printf("Scraped %d listings so far\n", len(listings))
	}

	log.Println(len(listings), " listings scraped successfully")
	return listings, nil
}
