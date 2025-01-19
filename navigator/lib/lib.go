package lib

import (
	"math"
)

var earthRadius = 6371.0 // Earth's radius in kilometers

// Function to calculate the great-circle distance
func calculateDistance(lat1, lon1, lat2, lon2, earthRadius float64) float64 {
	// Calculate the distance using the given formula
	distance := math.Acos(math.Sin(lat1)*math.Sin(lat2)+
		math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1)) * earthRadius

	return distance
}

func GetDistance(lat1Deg, lon1Deg, lat2Deg, lon2Deg float64) float64 {
	// Convert degrees to radians
	lat1 := degToRad(lat1Deg)
	lon1 := degToRad(lon1Deg)
	lat2 := degToRad(lat2Deg)
	lon2 := degToRad(lon2Deg)

	// Calculate the distance
	distance := calculateDistance(lat1, lon1, lat2, lon2, earthRadius)

	return distance
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
