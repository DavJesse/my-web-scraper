package main

import (
	"log"
	"my-web-scraper/store"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://jiji.co.ke/cars")
	var title string

	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	var records []string

	records = append(records, "Title")

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".b-list-advert-base__data").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
        records = append(records, title)
	})
	
	store.SaveToJSON(records)
}
