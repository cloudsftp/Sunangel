package location

import (
	"math"

	"github.com/cloudsftp/Sunangel/angle"
)

func (a Location) angleAtCenterOfEarth(b Location) float64 {
	longa := angle.RadiansFromDegrees(a.Longitude)
	longb := angle.RadiansFromDegrees(b.Longitude)

	lata := angle.RadiansFromDegrees(a.Latitude)
	latb := angle.RadiansFromDegrees(b.Latitude)

	Aa := math.Cos(longa) * math.Cos(lata)
	Ab := math.Cos(longb) * math.Cos(latb)

	Ba := math.Cos(longa) * math.Sin(lata)
	Bb := math.Cos(longb) * math.Sin(latb)

	Ca := math.Sin(longa)
	Cb := math.Sin(longb)

	result := Aa*Ab + Ba*Bb + Ca*Cb
	result = math.Acos(result)
	return angle.NormalizeRadians(result)
}
