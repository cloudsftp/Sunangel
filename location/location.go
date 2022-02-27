package location

type horizonArray [horizonAngleResolution]float64

type Location struct {
	Latitude  float64
	Longitude float64
	Horizon   horizonArray
}

func NewLocation(latitude, longitude float64) *Location {
	l := Location{Latitude: latitude, Longitude: longitude}
	return &l
}
