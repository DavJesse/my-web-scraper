package main

import (
	"log"

	"my-web-scraper/services"
	"my-web-scraper/store"
)

func main() {
	// Define key variables; target url and target div className
	path := "https://jiji.co.ke/mombasa-cbd/buses"
	targetDiv := ".b-list-advert-base__data__inner"

	// Scrape data from web page, handle errors if encountered
	listings, err := services.ScrapeDataWithHeadless(path, targetDiv)
	if err != nil {
		log.Fatal(err)
	}

	// Write scraped data to JSON file
	store.SaveToJSON(listings)
}
