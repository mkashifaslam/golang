package city

import (
	"errors"
	"fmt"
	"github.com/mkashifaslam/golang/navigator/csv"
	"github.com/mkashifaslam/golang/navigator/lib"
	"github.com/mkashifaslam/golang/navigator/location"
	"strings"
)

type City struct {
	Name     string
	Location location.Location
}

type Compass int

//go:generate stringer --type Compass --trimprefix compass --linecomment
const (
	_           Compass = iota
	Origin              // Origin
	Destination         // Destination
)

func New(name string, latD, lngD string) *City {
	lat, lng := lib.ConvToFloat64(latD), lib.ConvToFloat64(lngD)
	return &City{
		Name:     name,
		Location: *location.New(lat, lng),
	}
}

func (c *City) Print() {
	fmt.Printf("%s @ %+v\n", c.Name, c.Location)
}

func (c *City) GetName() string {
	return c.Name
}

func (c *City) GetLocation() location.Location {
	return c.Location
}

func (c *City) GetLat() float64 {
	return c.Location.Lat
}

func (c *City) GetLng() float64 {
	return c.Location.Lng
}

func (c *City) Distance(cityB City) float64 {
	return lib.GetDistance(c.GetLat(), c.GetLng(), cityB.GetLat(), cityB.GetLng())
}

func GetList(cities csv.Data) []City {
	var formattedCities []City
	for name, value := range cities {
		lat, lng := value["lat"], value["lng"]
		formattedCities = append(formattedCities, *New(name, lat, lng))
	}
	return formattedCities
}

func Search(cities []City, name string, direction Compass) (City, error) {
	for _, city := range cities {
		if compare(city.Name, name) {
			fmt.Printf("%s city found: %s\n", direction.String(), city.Name)
			return city, nil
		}
	}
	return City{}, errors.New("city not found")
}

func compare(storeCity, inputCity string) bool {
	var cityA, cityB = strings.ToLower(storeCity), strings.ToLower(inputCity)
	return strings.Contains(cityA, cityB)
}

func PrintAll(cities []City) {
	for _, city := range cities {
		city.Print()
	}
}
