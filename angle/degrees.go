package angle

import "math"

const (
	degreesPerRadian   float64 = 180 / math.Pi
	degreesPerRotation float64 = 360
)

func DegreesFromRadians(radians float64) float64 {
	return radians * degreesPerRadian
}

func NormalizeDegrees(degrees float64) float64 {
	degrees = math.Mod(degrees, degreesPerRotation)
	if degrees < 0 {
		degrees += degreesPerRotation
	}
	return degrees
}
