package location

import (
	"math"

	"github.com/cloudsftp/Sunangel/angle"
)

func (src Location) AltitudeAngleTo(tgt *Location) float64 {
	dist := src.distance(*tgt)

	heightSrc := src.GetElevation()
	heightTgt := tgt.GetElevation()

	alt := math.Atan2(heightTgt-heightSrc, dist)
	alt -= dist / (2 * radiusEarthMeters)

	return angle.NormalizeRadiansLatitude(alt)
}

/*

This procedure produces angles that are too small!

Considers the earth as round
Rounding errors probably result in wrong angles: theta is very small!

func (src Location) HorizontalAngleTo(tgt *Location) float64 {
	theta := src.angleAtCenterOfEarth(tgt)

	heightA := src.GetElevation()
	heightB := tgt.GetElevation()

	d1 := 2 * math.Sin(theta/2) * (radiusEarthMeters + heightA)

	counterThetaHalves := (math.Pi - theta) / 2
	dh := heightB - heightA
	h := math.Sin(counterThetaHalves) * dh
	d2 := math.Cos(counterThetaHalves) * dh

	d := d1 + d2

	gamma := math.Atan2(h, d)
	gamma = gamma - (theta / 2)
	gamma = angle.NormalizeRadiansLatitude(gamma)
	return gamma
}
*/

func (src Location) AzimutAngleTo(tgt *Location) float64 {
	srcLat := angle.RadiansFromDegrees(src.Latitude)
	tgtLat := angle.RadiansFromDegrees(tgt.Latitude)
	srcLong := angle.RadiansFromDegrees(src.Longitude)
	tgtLong := angle.RadiansFromDegrees(tgt.Longitude)

	dlong := tgtLong - srcLong

	y := math.Sin(dlong) * math.Cos(tgtLat)
	x := math.Cos(srcLat) * math.Sin(tgtLat)
	x -= math.Sin(srcLat) * math.Cos(tgtLat) * math.Cos(dlong)

	azimutAngle := math.Atan2(y, x)
	azimutAngle = angle.NormalizeRadians(azimutAngle)
	return azimutAngle
}
