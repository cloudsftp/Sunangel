package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/sunangel"
)

func TestAltitudeSunangleAt(t *testing.T) {
	got := sunangel.AltitudeSunAngleAt(dateWiki, latitudeWiki, longitudeWiki)
	want := float64(19.062)

	assertApproxEqual(t, got, want)
}
