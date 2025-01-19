package location

type Location struct {
	Lat float64
	Lng float64
}

func New(lat, lng float64) *Location {
	return &Location{
		Lat: lat,
		Lng: lng,
	}
}
