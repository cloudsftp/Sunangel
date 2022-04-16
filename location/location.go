package location

type Location struct {
	Latitude  float64
	Longitude float64
}

// NewLocation returns a new location struct from a name and coordinates.
func NewLocation(latitude, longitude float64) *Location {
	return &Location{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
