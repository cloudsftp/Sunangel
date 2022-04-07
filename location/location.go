package location

type horizonArray [horizonAngleResolution]float64

type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
	Horizon   horizonArray
}

// NewLocation returns a new location struct from a name and coordinates.
func NewLocation(name string, latitude, longitude float64) *Location {
	return &Location{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
