package main

import (
	"flycode.go/geocity/database"
	"flycode.go/geocity/generate"
)

func main() {
	database.Migrate()

	// No entry in database so generate all the database
	if database.CityCount() == 0 {
		generate.Generate("public/data/fr.csv")
		// Get all lat & long with Google Place engine
		//generate.LinkPlaces()
	}
	generate.LinkPlaces()

	// Get the postal code (better if is set)

	// Get the postal city

	// Return the result

}
