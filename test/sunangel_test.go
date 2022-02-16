package test

import (
	"fmt"
	"testing"

	"github.com/cloudsftp/Sunangel/angle"
	"github.com/cloudsftp/Sunangel/sunangel"
)

func TestAltitudeSunangleAt(t *testing.T) {
	got := sunangel.AltitudeSunAngleAt(dateWiki, latitudeWiki, longitudeWiki)
	want := angle.RadiansFromDegrees(19.062)

	assertApproxEqual(t, got, want)
}

func TestAzimutSunAngleAt(t *testing.T) {
	got := sunangel.AzimutSunAngleAt(dateWiki, latitudeWiki, longitudeWiki)
	want := angle.RadiansFromDegrees(265.938)

	assertApproxEqual(t, got, want)
}

func TestAltitudeSunangleAtCustom(t *testing.T) {
	got := sunangel.AltitudeSunAngleAt(dateCustom, latitudeCustom, longitudeCustom)

	fmt.Printf("got %.5f\n", angle.DegreesFromRadians(got))
}
