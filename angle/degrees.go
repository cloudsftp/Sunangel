package angle

import "math"

const (
	degreesPerRadian   float64 = 180 / math.Pi
	degreesPerRotation float64 = 360
)

// DegreesFromRadians returns the angle in degrees which is given in radians.
// It will be in the range [0, 360)
func DegreesFromRadians(radians float64) float64 {
	degrees := radians * degreesPerRadian
	return NormalizeDegrees(degrees)
}

// NormalizeDegrees resturns the normalized angle in degrees.
// It will be in the range [0, 360)
func NormalizeDegrees(degrees float64) float64 {
	degrees = math.Mod(degrees, degreesPerRotation)
	if degrees < 0 {
		degrees += degreesPerRotation
	}
	return degrees
}

// NormalizeDegrees resturns the normalized angle in degrees.
// It will be in the range [-180, 180)
func NormalizeDegreesLatitude(degrees float64) float64 {
	degrees = NormalizeDegrees(degrees)
	if degrees >= degreesPerRotation/2 {
		degrees -= degreesPerRotation
	}
	return degrees
}
