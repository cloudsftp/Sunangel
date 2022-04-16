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

type AltitudeArray [horizonAngleResolution]float64

type Horizon struct {
	Place       *location.Location
	startRadius int // radius to ignore when computing the horizon in meters
	altitude    AltitudeArray
}

// NewHorizon returns the computed horizon for a given location.
// The computation starts at the given startRadius
func NewHorizon(place *location.Location, startRadius int) *Horizon {
	horizon := &Horizon{
		Place:       place,
		startRadius: startRadius,
		altitude:    AltitudeArray{},
	}

	horizon.compute()
	return horizon
}

// NewHorizonWithAltitude returns a horizon with a given startRadius and altitude
func NewHorizonWithAltitude(place *location.Location, r int, a AltitudeArray) *Horizon {
	horizon := &Horizon{
		Place:       place,
		startRadius: r,
		altitude:    a,
	}

	return horizon
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

// GetStartRadius returns the start radius of the given horizon.
// Only use for testing!
func (hor Horizon) GetStartRadius() int {
	return hor.startRadius
}

// GetAltitudeArray returns the altitude array of the given horizon.
// Only use for testing!
func (hor Horizon) GetAltitudeArray() AltitudeArray {
	return hor.altitude
}
