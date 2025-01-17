package city

import (
	"errors"
	"fmt"
	"github.com/mkashifaslam/golang/navigator/csv"
	"github.com/mkashifaslam/golang/navigator/lib"
	"github.com/mkashifaslam/golang/navigator/location"
)

type City struct {
	Name     string
	Location location.Location
}

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

func Search(cities []City, name string) (City, error) {
	for _, city := range cities {
		if city.Name == name {
			return city, nil
		}
	}

	return City{}, errors.New("city not found")
}

func PrintAll(cities []City) {
	for _, city := range cities {
		city.Print()
	}
}
