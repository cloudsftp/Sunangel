package horizon

import (
	"math"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/util"
)

const (
	horizonAngleResolution int     = 1 << 10 // 1024
	horizonAngleWidth      float64 = 2 * math.Pi / float64(horizonAngleResolution)
)

type horizonArray [horizonAngleResolution]float64

type Horizon struct {
	Place        *location.Location
	radiusIgnore float64 // radius to ignore when computing the horizon in meters
	altitude     horizonArray
}

func NewHorizon(place *location.Location, radiusIgnore float64) *Horizon {
	horizon := Horizon{
		Place:        place,
		radiusIgnore: radiusIgnore,
		altitude:     horizonArray{},
	}

	horizon.compute()
	return &horizon
}

// GetHorizonAngleAt returns the altitude angle of the horizon
// of a given location at a given azimut angle.
func (horizon Horizon) GetAltitude(tgtAzimutAngle float64) float64 {
	leftIndex := int(tgtAzimutAngle / horizonAngleWidth)
	rightIndex := (leftIndex + 1) % horizonAngleResolution

	leftAzimutAngle := float64(leftIndex) * horizonAngleWidth
	rightAzimutAngle := float64(rightIndex) * horizonAngleWidth

	leftHorizonAngle := horizon.altitude[leftIndex]
	rightHorizonAngle := horizon.altitude[rightIndex]

	return util.LinInt(tgtAzimutAngle, leftAzimutAngle, leftHorizonAngle, rightAzimutAngle, rightHorizonAngle)
}
