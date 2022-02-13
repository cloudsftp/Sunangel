package startime

import (
	"math"
	"time"
)

const (
	starTimeC0             float64 = 6.697376
	starTimeC1             float64 = 2400.05134
	starTimeC2             float64 = 1.002738
	hoursInDay             float64 = 24
	starTimeDegreesPerHour float64 = 15 * (math.Pi / 180)
)

func starTimeAt(date time.Time) float64 {
	t0 := julianCenturiesSince2000ToMidnightOf(date)
	t := float64(date.Hour()) + float64(date.Minute())/60

	starTime := starTimeC0 + starTimeC1*t0 + starTimeC2*t
	starTime = math.Mod(starTime, hoursInDay)
	return starTime
}

func greenwichHourAngleOfSpringPointAt(date time.Time) float64 {
	thetaGh := starTimeAt(date)
	thetaG := thetaGh * starTimeDegreesPerHour
	return math.Mod(thetaG, 2*math.Pi)
}

func hourAngleOfSpringPointAt(date time.Time, longitude float64) float64 {
	thetaG := greenwichHourAngleOfSpringPointAt(date)
	theta := thetaG + longitude
	return math.Mod(theta, 2*math.Pi)
}

func HourAngleOfSunAt(date time.Time, longitude float64) float64 {
	date = date.UTC() // all exported functions have to make sure, dates are UTC
	longitude *= (math.Pi / 180)

	theta := hourAngleOfSpringPointAt(date, longitude)
	alpha := rightAscensionOfSunAt(date)

	tau := theta - alpha
	return math.Mod(tau, 2*math.Pi)
}
