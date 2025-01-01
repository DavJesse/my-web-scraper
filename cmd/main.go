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
	path := "https://jiji.co.ke/cars"
	targetDiv := ".b-list-advert-base__data__inner"
	
	
	store.SaveToJSON(listings)
}
