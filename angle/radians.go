package angle

import "math"

const (
	radiansPerDegree   float64 = math.Pi / 180
	radiansPerRotation float64 = 2 * math.Pi
)

// RadiansFromDegrees returns the angle in radians which is given in degrees.
// It will be in the range [0, 2π)
func RadiansFromDegrees(degrees float64) float64 {
	radians := degrees * radiansPerDegree
	return NormalizeRadians(radians)
}

// NormalizeRadians returns the normalized angle in radians.
// It will be in the range [0, 2π)
func NormalizeRadians(radians float64) float64 {
	radians = math.Mod(radians, radiansPerRotation)
	if radians < 0 {
		radians += 2 * math.Pi
	}
	return radians
}

// NormalizeRadiansLatitude returns the normalized angle in radians.
// It will be in the range [-π, π)
func NormalizeRadiansLatitude(radians float64) float64 {
	radians = NormalizeRadians(radians)
	if radians >= math.Pi {
		radians -= 2 * math.Pi
	}
	return radians
}
