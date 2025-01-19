package lib

import (
	"fmt"
	"github.com/mkashifaslam/golang/navigator/location"
	"strconv"
)

func ConvToFloat64(value string) float64 {
	float, _ := strconv.ParseFloat(value, 64)
	return float
}

func GetFormattedDistance(distance float64) string {
	return fmt.Sprintf("%.2f km", distance)
}

func GetLatLng(data map[string]string) *location.Location {
	lat, lng := ConvToFloat64(data["lat"]), ConvToFloat64(data["lng"])
	return location.New(lat, lng)
}
