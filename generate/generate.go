package generate

import (
	"context"
	"encoding/csv"
	"flycode.go/geocity/database"
	"fmt"
	"googlemaps.github.io/maps"
	"io"
	"os"
)

func Generate(filePath string) {
	fmt.Print("Generate the database with the csv file.\n")

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'

	db, _ := database.GetDataBase()
	defer db.Close()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		city := database.CityNew(line[1], line[2])

		if db.NewRecord(city) == true {
			db.Create(&city)
		}
	}

	fmt.Print("Generate database complete.\n")
}

func LinkPlaces() {
	db, _ := database.GetDataBase()
	defer db.Close()

	cities := []*database.City{}
	db.Where("lat NOT LIKE 0 AND long NOT LIKE 0").Find(&cities)
	cpt := 1
	for _, city := range cities {
		client, err := maps.NewClient(maps.WithAPIKey("YOUR API KEY"))
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			panic("failed to connect at Google CLoud")
		}

		r := &maps.GeocodingRequest{
			Address: city.PostalCode + " " + city.FullName,
		}

		resp, err := client.Geocode(context.Background(), r)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			panic("failed to find place")
		}

		fmt.Printf("(%v/%v)  ", cpt, len(cities))

		if len(resp) > 0 {
			city.Lat = resp[0].Geometry.Location.Lat
			city.Long = resp[0].Geometry.Location.Lng

			fmt.Printf("+ City : %v found at %v, %v\n", city, city.Lat, city.Long)

			db.Save(&city)
		} else {
			fmt.Printf("- City : %v not found.\n", city)
		}

		cpt += 1
	}
}
