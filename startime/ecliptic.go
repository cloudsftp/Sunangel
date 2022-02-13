package startime

import (
	"math"
	"time"
)

const (
	meanEclipticLengthOfSunC0 float64 = 280.46 * (math.Pi / 180)
	meanEclipticLengthOfSunC1 float64 = 0.9856474 * (math.Pi / 180)

	meanAnomalyOfSunC0 float64 = 357.528 * (math.Pi / 180)
	meanAnomalyOfSunC1 float64 = 0.9856003 * (math.Pi / 180)

	eclipticLengthOfSunC0 float64 = 1.915 * (math.Pi / 180)
	eclipticLengthOfSunC1 float64 = 0.01997 * (math.Pi / 180)

	skewOfEclipticC0 float64 = 23.439 * (math.Pi / 180)
	skewOfEclipticC1 float64 = 0.4e-6 * (math.Pi / 180)
)

func meanEclipticLengthOfSunAt(date time.Time) float64 {
	l := meanEclipticLengthOfSunC0
	l += meanEclipticLengthOfSunC1 * julianDaysSince2000At(date)
	return math.Mod(l, 2*math.Pi)
}

func meanAnomalyOfSunAt(date time.Time) float64 {
	g := meanAnomalyOfSunC0
	g += meanAnomalyOfSunC1 * julianDaysSince2000At(date)
	return +math.Mod(g, 2*math.Pi)
}

func eclipticLengthOfSunAt(date time.Time) float64 {
	l := meanEclipticLengthOfSunAt(date)
	g := meanAnomalyOfSunAt(date)

	lambda := l
	lambda += eclipticLengthOfSunC0 * math.Sin(g)
	lambda += eclipticLengthOfSunC1 * math.Sin(2*g)
	return math.Mod(lambda, 2*math.Pi)
}

func skewOfEclipticAt(date time.Time) float64 {
	epsilon := skewOfEclipticC0
	epsilon += skewOfEclipticC1 * julianDaysSince2000At(date)
	return math.Mod(epsilon, 2*math.Pi)
}
