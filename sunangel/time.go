package sunangel

import (
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	JD_2000                    float64 = 2451545
	JULIAN_DAYS_IN_CENTURY     float64 = 36525
	STAR_TIME_C0               float64 = 6.697376
	STARTIME_C1                float64 = 2400.05134
	STARTIME_C2                float64 = 1.002738
	STAR_TIME_DEGREES_PER_HOUR float64 = 15
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
	return (julian.TimeToJD(date) - JD_2000) / JULIAN_DAYS_IN_CENTURY
}

func starTimeHoursAt(date time.Time) float64 {
	t0 := julianCenturiesSince2000To(date)
	t := float64(date.Hour()) + float64(date.Minute())/60

	starTime := STAR_TIME_C0 + STARTIME_C1*t0 + STARTIME_C2*t
	starTime = math.Mod(starTime, 24)
	return starTime
}

func starTimeAt(date time.Time) float64 {
	return starTimeHoursAt(date) * STAR_TIME_DEGREES_PER_HOUR
}
