package sunangel

import (
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	jd2000              float64 = 2451545
	julianDaysInCentury float64 = 36525
)

// JulianDaysSince2000At returns the julian days
// since the beginning of the year 2000 until the given time.
func julianDaysSince2000At(date time.Time) float64 {
	jd := julian.TimeToJD(date) - jd2000
	return jd
}

// MidnightOf returns a time struct at midnight of a given time struct.
func midnightOf(date time.Time) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()
	location := date.Location()

	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

// JulianDaysSince2000AtToMidnightOf returns the julian days
// since the beginning of the year 2000 until midnight of day of the given time.
func julianDaysSince2000ToMidnightOf(date time.Time) float64 {
	date = midnightOf(date)
	jd0 := julianDaysSince2000At(date)
	return jd0
}

// JulianCenturiesSince2000AtToMidnightOf returns the julian centuries
// since the beginning of the year 2000 until midnight of day of the given time.
func julianCenturiesSince2000ToMidnightOf(date time.Time) float64 {
	jd0 := julianDaysSince2000ToMidnightOf(date)
	t0 := jd0 / julianDaysInCentury
	return t0
}

// TimeOfDayAsDecimal returns the time of the day of the given time as a decimal number.
func timeOfDayAsDecimal(date time.Time) float64 {
	timeOfDay := float64(date.Hour())
	timeOfDay += float64(date.Minute()) / 60
	timeOfDay += float64(date.Second()) / 3600
	timeOfDay += float64(date.Nanosecond()) / (3600 * 1e9)

	return timeOfDay
}
