package store

import (
	"encoding/json"
	"log"
	"os"
)

func SaveToJSON(records []string) {
	file, err := os.Create("scraped_data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, data := range records {
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", " ") // Sets-up pretty-print
		err = encoder.Encode(data)
	}

	if err != nil {
		log.Fatalf("Failed encoding data to JSON: %s", err)
	}
	log.Println("JSON writen succesfully")
}
