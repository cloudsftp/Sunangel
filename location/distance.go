package location

import (
	"math"

	"github.com/cloudsftp/Sunangel/angle"
)

const radiusEarthMeters float64 = 6371e3

func (src Location) distance(tgt Location) float64 {
	latSrc := angle.RadiansFromDegrees(src.Latitude)
	latTgt := angle.RadiansFromDegrees(tgt.Latitude)
	latD := latTgt - latSrc

	lngD := angle.RadiansFromDegrees(tgt.Longitude - src.Longitude)

	haversine := math.Pow(math.Sin(latD/2), 2)
	haversine += math.Cos(latSrc) * math.Cos(latTgt) * math.Pow(math.Sin(lngD/2), 2)

	dist := math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))
	dist *= 2 * radiusEarthMeters
	return dist
}
