package location

import (
	"math"

	"github.com/cloudsftp/Sunangel/angle"
)

func (src Location) angleAtCenterOfEarth(tgt Location) float64 {
	srcLat := angle.RadiansFromDegrees(src.Latitude)
	tgtLat := angle.RadiansFromDegrees(tgt.Latitude)
	srcLong := angle.RadiansFromDegrees(src.Longitude)
	tgtLong := angle.RadiansFromDegrees(tgt.Longitude)

	srcA := math.Cos(srcLong) * math.Cos(srcLat)
	tgtA := math.Cos(tgtLong) * math.Cos(tgtLat)

	srcB := math.Cos(srcLong) * math.Sin(srcLat)
	tgtB := math.Cos(tgtLong) * math.Sin(tgtLat)

	srcC := math.Sin(srcLong)
	tgtC := math.Sin(tgtLong)

	result := srcA*tgtA + srcB*tgtB + srcC*tgtC
	result = math.Acos(result)
	return angle.NormalizeRadians(result)
}
