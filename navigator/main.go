package main

import (
	"embed"
	"fmt"
	"github.com/mkashifaslam/golang/navigator/city"
	"github.com/mkashifaslam/golang/navigator/csv"
	"github.com/mkashifaslam/golang/navigator/lib"
	"github.com/mkashifaslam/golang/str2num/input"
	"time"
)

//go:embed data/cities.csv
var static embed.FS
var csvFile = "data/cities.csv"

func main() {
	fmt.Println("PAK Geo Navigator")
	fmt.Println("Loading cities...")
	cities, err := csv.LoadData(csvFile, static)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Loaded %d cities\n", len(cities))
	fmt.Println("http://localhost:8080/cities")
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
	fmt.Println("Shortest Distance:", lib.GetFormattedDistance(distance))
}
