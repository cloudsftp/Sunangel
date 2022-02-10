package sunangel

import (
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	JD_2000                float64 = 2451545
	JULIAN_DAYS_IN_CENTURY float64 = 36525
)

func julianCenturiesSince2000To(time time.Time) float64 {
	return (julian.TimeToJD(time) - JD_2000) / JULIAN_DAYS_IN_CENTURY
}

func julianCenturiesSince2000() float64 {
	return julianCenturiesSince2000To(time.Now())
}
