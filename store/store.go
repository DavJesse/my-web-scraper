package store

import (
	"encoding/json"
	"log"
	"os"
)

func SaveToJSON(data interface{}) {
	// Marshal data to JSON with indentetion
	file, err := json.MarshalIndent(data, "", "	 ")
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	// Write JSON to a file with specified permissions and mode 0644 (readable by others)
	err = os.WriteFile("scraped_data.json", file, 0644)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
	}
	log.Println("JSON writen succesfully")
}
