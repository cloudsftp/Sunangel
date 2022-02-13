package startime

import (
	"math"
	"time"

	"github.com/cloudsftp/Sunangel/angle"
)

var (
	meanEclipticLengthOfSunC0 float64 = angle.RadiansFromDegrees(280.46)
	meanEclipticLengthOfSunC1 float64 = angle.RadiansFromDegrees(0.9856474)

	meanAnomalyOfSunC0 float64 = angle.RadiansFromDegrees(357.528)
	meanAnomalyOfSunC1 float64 = angle.RadiansFromDegrees(0.9856003)

	eclipticLengthOfSunC0 float64 = angle.RadiansFromDegrees(1.915)
	eclipticLengthOfSunC1 float64 = angle.RadiansFromDegrees(0.01997)

	skewOfEclipticC0 float64 = angle.RadiansFromDegrees(23.439)
	skewOfEclipticC1 float64 = angle.RadiansFromDegrees(0.4e-6)
)

func meanEclipticLengthOfSunAt(date time.Time) float64 {
	l := meanEclipticLengthOfSunC0
	l += meanEclipticLengthOfSunC1 * julianDaysSince2000At(date)
	return angle.NormalizeRadians(l)
}

func meanAnomalyOfSunAt(date time.Time) float64 {
	g := meanAnomalyOfSunC0
	g += meanAnomalyOfSunC1 * julianDaysSince2000At(date)
	return angle.NormalizeRadians(g)
}

func eclipticLengthOfSunAt(date time.Time) float64 {
	l := meanEclipticLengthOfSunAt(date)
	g := meanAnomalyOfSunAt(date)

	lambda := l
	lambda += eclipticLengthOfSunC0 * math.Sin(g)
	lambda += eclipticLengthOfSunC1 * math.Sin(2*g)
	return angle.NormalizeRadians(lambda)
}

func skewOfEclipticAt(date time.Time) float64 {
	epsilon := skewOfEclipticC0
	epsilon += skewOfEclipticC1 * julianDaysSince2000At(date)
	return angle.NormalizeRadians(epsilon)
}
