package services

import (
	"fmt"
	"net/http"
	"strings"

	"my-web-scraper/models"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeData(path, targetDiv string) ([]models.CarListing, error) {
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
	return listings, nil
}
