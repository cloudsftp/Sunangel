package startime

import (
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	jd2000                 float64 = 2451545
	julianDaysInCentury    float64 = 36525
	starTimeC0             float64 = 6.697376
	starTimeC1             float64 = 2400.05134
	starTimeC2             float64 = 1.002738
	hoursInDay             float64 = 24
	starTimeDegreesPerHour float64 = 15
)

// for all unexportet functions: date must be in UTC

func midnightOf(date time.Time) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func julianCenturiesSince2000To(date time.Time) float64 {
	date = midnightOf(date)
	return (julian.TimeToJD(date) - jd2000) / julianDaysInCentury
}

func starTimeAt(date time.Time) float64 {
	t0 := julianCenturiesSince2000To(date)
	t := float64(date.Hour()) + float64(date.Minute())/60

	starTime := starTimeC0 + starTimeC1*t0 + starTimeC2*t
	starTime = math.Mod(starTime, hoursInDay)
	return starTime
}

func greenwichHourAngleAt(date time.Time) float64 {
	return starTimeAt(date) * starTimeDegreesPerHour
}

func hourAngleAt(date time.Time, longitude float64) float64 {
	return greenwichHourAngleAt(date) + longitude
}

func HourAngleAt(date time.Time, longitude float64) float64 {
	return hourAngleAt(date.UTC(), longitude)
}

func JulianDaysSince2000At(date time.Time) float64 {
	return julian.TimeToJD(date.UTC()) - jd2000
}
