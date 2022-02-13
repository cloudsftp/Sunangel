package angle

import "math"

const (
	radiansPerDegree   float64 = math.Pi / 180
	radiansPerRotation float64 = 2 * math.Pi
)

func RadiansFromDegrees(degrees float64) float64 {
	return degrees * radiansPerDegree
}

func NormalizeRadians(radians float64) float64 {
	radians = math.Mod(radians, radiansPerRotation)
	if radians < 0 {
		radians += radiansPerRotation
	}
	return radians
}
