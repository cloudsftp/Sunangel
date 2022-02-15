package startime

import (
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	jd2000              float64 = 2451545
	julianDaysInCentury float64 = 36525
)

func JulianDaysSince2000At(date time.Time) float64 {
	jd := julian.TimeToJD(date) - jd2000
	return jd
}

func midnightOf(date time.Time) time.Time {
	year := date.Year()
	month := date.Month()
	day := date.Day()
	location := date.Location()

	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

func JulianDaysSince2000ToMidnightOf(date time.Time) float64 {
	date = midnightOf(date)
	jd0 := JulianDaysSince2000At(date)
	return jd0
}

func JulianCenturiesSince2000ToMidnightOf(date time.Time) float64 {
	jd0 := JulianDaysSince2000ToMidnightOf(date)
	t0 := jd0 / julianDaysInCentury
	return t0
}
