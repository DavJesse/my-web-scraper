package main

import (
	"log"

	"my-web-scraper/services"
	"my-web-scraper/store"
)

func main() {
	// Define the URL and target div for scraping
	path := "https://jiji.co.ke/nairobi/cars"
	targetDiv := ".b-list-advert-base__data__inner"

	// Scrape car listings, handle necessary errors
	listings, err := services.ScrapeData(path, targetDiv)
	if err != nil {
		log.Fatalf("Failed to scrape data from %s: %v.", path, err)
	}

	// Write scraped data to JSON file
	store.SaveToJSON(listings)
}
