package startime

import (
	"math"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
)

func rightAscensionOfSunAt(date time.Time) float64 {
	epsilon := skewOfEclipticAt(date)
	lambda := eclipticLengthOfSunAt(date)

	alpha := math.Atan(math.Cos(epsilon) * math.Tan(lambda))
	if math.Cos(lambda) < 0 {
		alpha += math.Pi
	}
	return angle.NormalizeRadians(alpha)
}

func declinationOfSunAt(date time.Time) float64 {
	epsilon := skewOfEclipticAt(date)
	lambda := eclipticLengthOfSunAt(date)

	delta := math.Asin(math.Sin(epsilon) * math.Sin(lambda))
	return angle.NormalizeRadians(delta)
}
