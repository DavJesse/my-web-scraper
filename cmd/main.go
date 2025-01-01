package main

import (
	"log"
	"my-web-scraper/models"
	"my-web-scraper/store"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var listings []models.CarListing
	response, err := http.Get("https://jiji.co.ke/cars")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".b-list-advert-base__data").Each(func(i int, s *goquery.Selection) {
		listing := models.CarListing{}

		// Extract title, Price, Description, and Location
		listing.Title = strings.TrimSpace(s.Find("b-advert-title-inner qa-advert-title b-advert-title-inner--div").Text())
		listing.Price = strings.TrimSpace(s.Find(".qa-advert-price").Text())
		listing.Description = strings.TrimSpace(s.Find(".b-list-advert-base__description-text").Text())
		listing.Location = strings.TrimSpace(s.Find(".b-list-advert__region__text").Text())
		listing.Condition = strings.TrimSpace(s.Find(".b-list-advert-base__item-attr").Text())

		listings = append(listings, listing)
	})
	
	store.SaveToJSON(listings)
}
