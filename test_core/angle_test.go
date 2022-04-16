package test_core

import (
	"math"
	"testing"

	"github.com/cloudsftp/Sunangel/angle"
)

func testRadiansFromDegreesGeneral(t *testing.T, degrees, radians float64) {
	got := angle.RadiansFromDegrees(degrees)
	assertPreciselyEqual(t, got, radians)
}

func TestRadiansFromDegrees(t *testing.T) {
	testRadiansFromDegreesGeneral(t, 0, 0)
	testRadiansFromDegreesGeneral(t, 90, math.Pi/2)
	testRadiansFromDegreesGeneral(t, 180, math.Pi)
	testRadiansFromDegreesGeneral(t, 360, 0)
}

func testNormalizeRadiansGeneral(t *testing.T, in, out float64) {
	got := angle.NormalizeRadians(in)
	assertPreciselyEqual(t, got, out)
}

func TestNormalizeRadians(t *testing.T) {
	testNormalizeRadiansGeneral(t, 0, 0)

	testNormalizeRadiansGeneral(t, 4.5*math.Pi, math.Pi/2)
	testNormalizeRadiansGeneral(t, 3*math.Pi, math.Pi)
	testNormalizeRadiansGeneral(t, 111*math.Pi, math.Pi)
	testNormalizeRadiansGeneral(t, 1e3*math.Pi, 0) // limit of this implementation

	testNormalizeRadiansGeneral(t, -2*math.Pi, 0)
}

func testDegreesFromRadiansGeneral(t *testing.T, degrees, radians float64) {
	got := angle.DegreesFromRadians(radians)
	assertPreciselyEqual(t, got, degrees)
}

func TestDegreesFromRadians(t *testing.T) {
	testDegreesFromRadiansGeneral(t, 0, 0)
	testDegreesFromRadiansGeneral(t, 90, math.Pi/2)
	testDegreesFromRadiansGeneral(t, 180, math.Pi)
	testDegreesFromRadiansGeneral(t, 0, 2*math.Pi)
}

func testNormalizeDegreesGeneral(t *testing.T, in, out float64) {
	got := angle.NormalizeDegrees(in)
	assertPreciselyEqual(t, got, out)
}

func TestNormalizeDegrees(t *testing.T) {
	testNormalizeDegreesGeneral(t, 0, 0)

	testNormalizeDegreesGeneral(t, 180, 180)
	testNormalizeDegreesGeneral(t, 111*180, 180)
	testNormalizeDegreesGeneral(t, 1e3*180, 0) // limit of this implementation

	testNormalizeDegreesGeneral(t, -180, 180)
}
