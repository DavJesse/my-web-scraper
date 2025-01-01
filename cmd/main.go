package main

import (
	"log"
	"strings"
	"time"

	"my-web-scraper/models"
	"my-web-scraper/services"
	"my-web-scraper/store"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
	path := "https://jiji.co.ke/mombasa-cbd/buses"
	targetDiv := ".b-list-advert-base__data__inner"
	
	store.SaveToJSON(listings)
}
