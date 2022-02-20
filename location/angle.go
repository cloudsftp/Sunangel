package location

import "math"

const circumferenceEarthMeters float64 = 6371e3

func (a Location) AngleTo(b Location) float64 {
	theta := a.angleAtCenterOfEarth(b)

	heightA := a.GetElevation()
	heightB := b.GetElevation()

	d1 := 2 * math.Sin(theta/2) * (circumferenceEarthMeters + heightA)

	counterThetaHalves := (math.Pi - theta) / 2
	dh := heightB - heightA
	h := math.Sin(counterThetaHalves) * dh
	d2 := math.Cos(counterThetaHalves) * dh

	d := d1 + d2

	gamma := math.Atan2(h, d)
	return gamma - (theta / 2)
}
