package main

import (
	"encoding/json"
	"flag"
	"flycode.go/geocity/database"
	"flycode.go/geocity/generate"
	"fmt"
	"os"
)

func main() {
	database.Migrate()

	command := os.Args[1:]

	if len(command) == 0 {
		panic("No command send. Do 'geocity help' for list command")
	}

	// Command line parsing
	switch command[0] {
	case "build": //geocity build
		// No entry in database so generate all the database
		if database.CityCount() == 0 {
			generate.Generate("public/data/fr.csv")
			// Get all lat & long with Google Place engine
			generate.LinkPlaces()
		} else {
			panic("You already build your data. For reset your data use the reset command")
		}
	case "reset": //geocity reset
		// No entry in database so generate all the database
		if database.CityCount() != 0 {
			os.Remove("database.db")
			database.Migrate()
		}
		generate.Generate("public/data/fr.csv")
		// Get all lat & long with Google Place engine
		generate.LinkPlaces()
	case "search": //geocity search --cp "55100" --city "verdun"
		myFlag := flag.NewFlagSet("", flag.ExitOnError)
		cp := myFlag.String("cp", "", "Postal Code")
		name := myFlag.String("city", "", "City Name")
		myFlag.Parse(os.Args[2:])

		cities := database.CitySearch(*cp, *name)
		for _, city := range cities {
			fmt.Printf("  - %v %v {%v, %v}\n", city.PostalCode, city.FullName, city.Lat, city.Long)
		}
	case "export": //geocity export --cp "55100" --city "verdun" > result.json
		myFlag := flag.NewFlagSet("", flag.ExitOnError)
		cp := myFlag.String("cp", "", "Postal Code")
		name := myFlag.String("city", "", "City Name")
		myFlag.Parse(os.Args[2:])

		cities := database.CitySearch(*cp, *name)
		b, err := json.Marshal(cities)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	default:
		panic("Unknown command send. Do 'geocity help' for list command")
	}
}
