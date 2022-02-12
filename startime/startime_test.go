package startime

import (
	"testing"

	"github.com/cloudsftp/Sunangel/util"
)

func TestJulianCenturiesSince2000ToMidnightOf(t *testing.T) {
	got := julianCenturiesSince2000ToMidnightOf(util.DateWiki)

	jd := float64(2453953.5)
	want := float64((jd - 2451545) / 36525)

	util.AssertApproxEqual(t, got, want)
}

func TestJulianCenturiesSince2000ToMidnightOfCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := julianCenturiesSince2000ToMidnightOf(util.DateCustom.UTC())
	want := float64(0.221122518823)

	util.AssertPreciselyEqual(t, got, want)
}

func TestStarTimeAt(t *testing.T) {
	got := starTimeAt(util.DateWiki)
	want := float64(2.9759)

	util.AssertApproxEqual(t, got, want)
}

func TestStartTimeAtCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := starTimeAt(util.DateCustom.UTC())
	want := float64(1.947950604654)

	util.AssertPreciselyEqual(t, got, want)
}

func TestGreenwichHourAngleAt(t *testing.T) {
	got := greenwichHourAngleAt(util.DateWiki)
	want := float64(44.6387)

	util.AssertApproxEqual(t, got, want)
}

func TestGreenwichHourAngleAtCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := greenwichHourAngleAt(util.DateCustom.UTC())
	want := float64(29.219259069814)

	util.AssertPreciselyEqual(t, got, want)
}

func TestHourAngleAtPrivate(t *testing.T) {
	got := hourAngleAt(util.DateWiki, float64(11.6))
	want := float64(56.2387)

	util.AssertApproxEqual(t, got, want)
}

func TestHourAngleAtCustomPrivate(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := hourAngleAt(util.DateCustom.UTC(), float64(9.58675))
	want := float64(38.806009069814)

	util.AssertPreciselyEqual(t, got, want)
}
