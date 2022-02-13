package startime

import (
	"math"
	"time"
)

func rightAscensionOfSunAt(date time.Time) float64 {
	epsilon := skewOfEclipticAt(date)
	lambda := eclipticLengthOfSunAt(date)

	alpha := math.Atan(math.Cos(epsilon) * math.Tan(lambda))
	if math.Cos(lambda) < 0 {
		alpha += math.Pi
	}
	return math.Mod(alpha, 2*math.Pi)
}

func declinationOfSunAt(date time.Time) float64 {
	epsilon := skewOfEclipticAt(date)
	lambda := eclipticLengthOfSunAt(date)

	delta := math.Asin(math.Sin(epsilon) * math.Sin(lambda))
	return math.Mod(delta, 2*math.Pi)
}
