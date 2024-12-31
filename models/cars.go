package models

type CarListing struct {
	Title       string `json:"title"`
	Price       string	`json:"price"`
	Description string `json:"description"`
	Condition string `json:"condition"`
}