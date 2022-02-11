package sunangel

import (
	"testing"

	"github.com/cloudsftp/Sunangel/test"
)

func TestJulianCenturiesSince2000To(t *testing.T) {
	got := julianCenturiesSince2000To(test.DateWiki)

	jd := float64(2453953.5)
	want := float64((jd - 2451545) / 36525)

	test.AssertApproxEqual(t, got, want)
}

func TestJulianCenturiesSince2000ToCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := julianCenturiesSince2000To(test.DateCustom.UTC())
	want := float64(0.221122518823)

	test.AssertPreciselyEqual(t, got, want)
}

func TestStarTimeHoursAt(t *testing.T) {
	got := starTimeHoursAt(test.DateWiki)
	want := float64(2.9759)

	test.AssertApproxEqual(t, got, want)
}

func TestStartTimeHoursAtCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := starTimeHoursAt(test.DateCustom.UTC())
	want := float64(1.947950604654)

	test.AssertPreciselyEqual(t, got, want)
}

func TestStarTimeAt(t *testing.T) {
	got := starTimeAt(test.DateWiki)
	want := float64(44.6387)

	test.AssertApproxEqual(t, got, want)
}

func TestStartTimeAtCustom(t *testing.T) {
	// date has to be in UTC for unexportet functions
	got := starTimeAt(test.DateCustom.UTC())
	want := float64(29.219259069814)

	test.AssertPreciselyEqual(t, got, want)
}
