package sunangel

import (
	"testing"
	"time"
)

func TestJulianCenturiesSince2000To(t *testing.T) {
	time := time.Date(2006, time.August, 6, 6, 0, 0, 0, time.UTC)
	got := julianCenturiesSince2000To(time)

	jd := float64(2453953.75)
	want := float64((jd - 2451545) / 36525)

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
