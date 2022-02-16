package sunangel

import (
	"math"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/location"
)

// AltitudeSunAngleAt returns the altitude of the sun at a given time and place.
// The latitude is the vertical angle between the sun and the horizon.
func AltitudeSunAngleAt(date time.Time, place location.Location) float64 {
	delta := declinationOfSunAt(date)
	tau := hourAngleOfSunAt(date, place.Longitude)

	argument := math.Cos(delta) * math.Cos(tau) * math.Cos(place.Latitude)
	argument += math.Sin(delta) * math.Sin(place.Latitude)

	result := math.Asin(argument)
	return angle.NormalizeRadians(result)
}

// AzimutSunAngleAt returns the azimut of the sun at a given time and place.
// The azimut is the horizontal angle between the sun to the orientation north.
func AzimutSunAngleAt(date time.Time, place location.Location) float64 {
	delta := declinationOfSunAt(date)
	tau := hourAngleOfSunAt(date, place.Longitude)

	nominator := math.Sin(tau)
	denominator := math.Cos(tau) * math.Sin(place.Latitude)
	denominator -= math.Tan(delta) * math.Cos(place.Latitude)

	argument := nominator / denominator
	result := math.Atan(argument)
	if denominator < 0 {
		result += math.Pi
	}
	return angle.NormalizeRadians(result)
}
