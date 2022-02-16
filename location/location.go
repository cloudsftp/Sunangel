package location

type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(latitude, longitude float64) *Location {
	l := Location{Latitude: latitude, Longitude: longitude}
	return &l
}
