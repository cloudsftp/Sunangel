package horizon

import (
	"log"
	"math"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/util"
)

const (
	stepSize        float64 = 0.0003
	stepSizeMeters  float64 = 30      // approximately
	maxSampleRadius int     = 2 << 10 // 1024
)

func (horizon *Horizon) compute() {
	log.Printf("Computing horizon for location %f, %f\n", horizon.Place.Latitude, horizon.Place.Longitude)

	for i := 0; i < len(horizon.altitude); i++ {
		horizon.altitude[i] = -math.Pi
	}

	for sampleRadius := 1; sampleRadius <= maxSampleRadius; sampleRadius++ {
		azimutAnglesMeasured, horizonAnglesMeasured := horizon.measureHorizonAngles(sampleRadius)

		currHorizonAngleResolution := computeSampleResolution(len(horizonAnglesMeasured))
		horizon.interpolateHorizonAnglesFromSamples(azimutAnglesMeasured, horizonAnglesMeasured, currHorizonAngleResolution)
		if currHorizonAngleResolution < horizonAngleResolution {
			horizon.interpolateHorizonAnglesFromHorizonAngles(currHorizonAngleResolution)
		}
	}
}

func computeSampleResolution(measurementResolution int) int {
	sampleResolution := math.Pow(2, math.Floor(math.Log2(float64(measurementResolution))))
	return int(math.Min(sampleResolution, float64(horizonAngleResolution)))
}

func (horizon Horizon) measureHorizonAngles(sampleRadius int) ([]float64, []float64) {
	offsets := circle(sampleRadius)

	skips := int(math.Max(float64(len(offsets)/(2*horizonAngleResolution)), 1))
	numSamples := int(len(offsets) / skips)

	azimutAngles := make([]float64, numSamples)
	horizonAngles := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		sampleLocation := computeOffsetLocation(horizon.Place, offsets[i*skips])

		azimutAngles[i] = horizon.Place.AzimutAngleTo(sampleLocation)
		horizonAngles[i] = horizon.Place.AltitudeAngleTo(sampleLocation)
	}

	return azimutAngles, horizonAngles
}

func computeOffsetLocation(place *location.Location, offset offsetCoordinates) *location.Location {
	offsetLocation := location.Location{}
	offsetLocation.Latitude = place.Latitude + float64(offset.latitude)*stepSize
	offsetLocation.Longitude = place.Longitude + float64(offset.longitude)*stepSize

	return &offsetLocation
}

func (horizon *Horizon) interpolateHorizonAnglesFromSamples(azimutAngles, horizonAngles []float64, resolution int) {
	sampleIndex := 0
	skips := horizonAngleResolution / resolution // resolution always power of 2 and < horizonAngleResolution

	for i := 0; i < resolution; i++ {
		targetAzimutAngle := float64(i*2) * math.Pi / float64(resolution)

		for ; sampleIndex < len(azimutAngles); sampleIndex++ {
			if azimutAngles[sampleIndex] > targetAzimutAngle {
				sampleIndex--
				break
			}
		}
		if sampleIndex >= len(azimutAngles) {
			sampleIndex = len(azimutAngles) - 1
		}

		leftAzimutAngle := azimutAngles[sampleIndex]
		leftHorizonAngle := horizonAngles[sampleIndex]

		rightIndex := (sampleIndex + 1) % len(azimutAngles)
		rightAzimutAngle := azimutAngles[rightIndex]
		rightHorizonAngle := horizonAngles[rightIndex]

		horizonAngle := util.LinInt(targetAzimutAngle, leftAzimutAngle, leftHorizonAngle, rightAzimutAngle, rightHorizonAngle)
		if horizonAngle > horizon.altitude[i*skips] {
			horizon.altitude[i*skips] = horizonAngle
		}
	}
}

func (horizon *Horizon) interpolateHorizonAnglesFromHorizonAngles(resolution int) {
	skipsInterpolated := horizonAngleResolution / resolution // resolution always power of 2 and < horizonAngleResolution
	offsetNotInterpolated := skipsInterpolated / 2

	for i := 0; i < resolution; i++ {
		leftIndex := i * skipsInterpolated
		rightIndex := (leftIndex + skipsInterpolated) % len(horizon.altitude)
		tgtIndex := leftIndex + offsetNotInterpolated

		leftHorizonAngle := horizon.altitude[leftIndex]
		rightHorizonAngle := horizon.altitude[rightIndex]

		horizonAngle := util.LinInt(1, 0, leftHorizonAngle, 2, rightHorizonAngle)
		if horizonAngle > horizon.altitude[tgtIndex] {
			horizon.altitude[tgtIndex] = horizonAngle
		}
	}
}
