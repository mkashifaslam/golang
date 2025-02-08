package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/navigator/city"
	"github.com/mkashifaslam/golang/navigator/csv"
	"github.com/mkashifaslam/golang/navigator/lib"
	"github.com/mkashifaslam/golang/str2num/input"
	"time"
)

var csvFile = "data/cities.csv"

func main() {
	fmt.Println("PAK Geo Navigator")
	fmt.Println("Loading cities...")
	cities, err := csv.LoadData(csvFile)
	if err != nil {
		_ = fmt.Errorf("error loading cities.csv: %v", err)
		return
	}

	fmt.Printf("Loaded %d cities\n", len(cities))
	fmt.Println("http://localhost:8080")
	go RunServer()
	formattedCities := city.GetList(cities)

	//city.PrintAll(formattedCities)

	origin := "Kāmoke, Gujranwala, Punjab, Pakistan"
	destination := "Pir Mahal, Toba Tek Singh, Punjab, Pakistan"

	fmt.Printf("\n--------------------------------------\n")

	origin = input.GetStringInput("Enter origin city name:")

	fmt.Println("Searching origin city...")
	time.Sleep(time.Second * 2)
	cityA, err := city.Search(formattedCities, origin, city.Origin)
	if err != nil {
		fmt.Println(err)
		return
	}

	destination = input.GetStringInput("Enter destination city name:")
	fmt.Println("Searching destination city...")
	time.Sleep(time.Second * 2)
	cityB, err := city.Search(formattedCities, destination, city.Destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	distance := cityA.Distance(cityB)
	fmt.Println("Distance:", lib.GetFormattedDistance(distance))
}
