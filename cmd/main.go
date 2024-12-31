package main

import (
	"log"
	"strings"
	"time"

	"my-web-scraper/models"
	"my-web-scraper/services"
	"my-web-scraper/store"

	"github.com/tebeka/selenium"
)

func main() {
	// Launch headless browser and create a new remote client instance
	wd, err := services.LaunchHeadlessBrowser()
	if err != nil {
		log.Fatal("Failed to Create New Remote Client: ", err)
	}
	defer wd.Quit()

	var listings []models.CarListing

	// Navigate to web page
	err = wd.Get("https://jiji.co.ke/mombasa-cbd/cars")
	if err != nil {
		log.Fatal(err)
		return
	}

	for len(listings) < 100 {
		// Scroll to bottom of the page to load more
		if _, err := wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil); err != nil {
			log.Fatal("Failed to Execute Scroll Script: ", err)
		}

		// Wait for contents to load
		time.Sleep(time.Second * 2)

		// Find all car listings on the page
		elements, err := wd.FindElements(selenium.ByCSSSelector, ".b-list-advert-base__data")
		if err != nil {
			log.Fatal("Failed to Find Elements: ", err)
		}

		newListings := len(elements)
		// break loop when no new content
		if newListings <= len(listings) {
			break
		}

		// Scrape listings
		for i := len(listings); i < newListings; i++ {
			element := elements[i]
			listing := models.CarListing{}

			if title, err := element.FindElement(selenium.ByCSSSelector, ".b-advert-title-inner.qa-advert-title.b-advert-title-inner--div"); err == nil {
				listing.Title, _ = title.Text()
			}
			if price, err := element.FindElement(selenium.ByCSSSelector, ".qa-advert-price"); err == nil {
				listing.Price, _ = price.Text()
			}
			if description, err := element.FindElement(selenium.ByCSSSelector, ".b-list-advert-base__description-text"); err == nil {
				listing.Description, _ = description.Text()
			}
			if location, err := element.FindElement(selenium.ByCSSSelector, ".b-list-advert__region__text"); err == nil {
				listing.Location, _ = location.Text()
			}
			if condition, err := element.FindElement(selenium.ByCSSSelector, ".b-list-advert-base__item-attr"); err == nil {
				listing.Condition, _ = condition.Text()
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
	store.SaveToJSON(listings)
}
