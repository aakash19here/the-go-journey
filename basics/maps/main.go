package main

import (
	"fmt"
	"log"
)

func main() {
	capitalCities := map[string]string{
		"United Arab Emirates": " Abu Dhabi",
		"Japan":                "Tokyo",
	}

	_, exists := capitalCities["Germany"]

	if !exists {
		log.Fatal("Country does not exist in the map")
	}

	for key, value := range capitalCities {
		fmt.Println("Capital of", key, "is", value)
	}
}
