package test

import (
	"testing"

	"github.com/cloudsftp/Sunangel/sunangel"
)

func TestSunangleAt(t *testing.T) {
	got := sunangel.SunAngleAt(dateWiki, latitudeWiki, longitudeWiki)
	want := float64(19.062)

	assertApproxEqual(t, got, want)
}
