package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"my-web-scraper/models"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func ScrapeData(path, targetDiv string) ([]models.CarListing, error) {
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
