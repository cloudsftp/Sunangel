package sunangel

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
	return alpha
}

func declinationOfSunAt(date time.Time) float64 {
	epsilon := skewOfEclipticAt(date)
	lambda := eclipticLengthOfSunAt(date)

	return math.Asin(math.Sin(epsilon) * math.Sin(lambda))
}
