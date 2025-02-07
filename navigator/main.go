package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/navigator/city"
	"github.com/mkashifaslam/golang/navigator/csv"
	"github.com/mkashifaslam/golang/navigator/lib"
	"github.com/mkashifaslam/golang/str2num/input"
)

var csvFile = "cities.csv"

func main() {
	fmt.Println("PAK Geo Navigator")
	fmt.Println("Loading cities...")
	cities, err := csv.LoadData(csvFile)
	if err != nil {
		_ = fmt.Errorf("error loading cities.csv: %v", err)
		return
	}

	fmt.Printf("Loaded %d cities\n", len(cities))

	formattedCities := city.GetList(cities)

	//city.PrintAll(formattedCities)

	origin := "Kāmoke, Gujranwala, Punjab, Pakistan"
	destination := "Pir Mahal, Toba Tek Singh, Punjab, Pakistan"

	fmt.Printf("\n--------------------------------------\n")

	origin = input.GetStringInput("Enter origin city name:")
	destination = input.GetStringInput("Enter destination city name:")

	fmt.Println("Searching origin city...")
	cityA, err := city.Search(formattedCities, origin)
	if err != nil {
		fmt.Printf("error searching cities: %v", err)
		return
	}
	fmt.Println("Origin city found")
	fmt.Println("Searching destination city...")
	cityB, err := city.Search(formattedCities, destination)
	if err != nil {
		fmt.Printf("error searching cities: %v", err)
		return
	}
	fmt.Println("Destination city found")
	distance := cityA.Distance(cityB)
	fmt.Println("Distance:", lib.GetFormattedDistance(distance))
}
