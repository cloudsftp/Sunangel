package location

import "math"

const circumferenceEarthMeters float64 = 6371e3

func (src Location) HorizontalAngleTo(tgt Location) float64 {
	theta := src.angleAtCenterOfEarth(tgt)

	heightA := src.GetElevation()
	heightB := tgt.GetElevation()

	d1 := 2 * math.Sin(theta/2) * (circumferenceEarthMeters + heightA)

	counterThetaHalves := (math.Pi - theta) / 2
	dh := heightB - heightA
	h := math.Sin(counterThetaHalves) * dh
	d2 := math.Cos(counterThetaHalves) * dh

	d := d1 + d2

	gamma := math.Atan2(h, d)
	return gamma - (theta / 2)
}

func (src Location) azimutAngleTo(tgt Location) float64 {
	dlong := tgt.Longitude - src.Longitude

	y := math.Sin(dlong) * math.Cos(tgt.Latitude)
	x := math.Cos(src.Latitude) * math.Sin(tgt.Latitude)
	x -= math.Sin(src.Latitude) * math.Cos(tgt.Latitude) * math.Cos(dlong)

	return math.Atan2(y, x)
}
