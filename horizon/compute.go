package horizon

import (
	"log"
	"math"

	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/util"
)

const (
	stepSize float64 = 0.0003
	numSteps int     = 2 << 11 // 2048
)

func (horizon *Horizon) compute() {
	log.Printf("Computing horizon for location %f, %f\n", horizon.Place.Latitude, horizon.Place.Longitude)

	for i := 0; i < len(horizon.altitude); i++ {
		horizon.altitude[i] = -math.Pi
	}

	for k := 16; k <= numSteps; k++ {
		azimutAnglesMeasured, horizonAnglesMeasured := horizon.measureHorizonAngles(k)

		currHorizonAngleResolution := computeSampleResolution(len(horizonAnglesMeasured))
		horizon.interpolateHorizonAnglesFromSamples(azimutAnglesMeasured, horizonAnglesMeasured, currHorizonAngleResolution)
		if currHorizonAngleResolution < horizonAngleResolution {
			horizon.interpolateHorizonAnglesFromHorizonAngles(currHorizonAngleResolution)
		}
	}
}

func (horizon Horizon) measureHorizonAngles(distance int) ([]float64, []float64) {
	skips := horizon.computeSkips(distance)
	numSamplesPerDiagonal := int(math.Ceil(float64(distance) / float64(skips)))
	numSamples := 4 * numSamplesPerDiagonal

	azimutAngles := make([]float64, numSamples)
	horizonAngles := make([]float64, numSamples)

	for i := 0; i < numSamplesPerDiagonal; i++ {
		offset := i * skips
		horizon.measureHorizonAngle(distance-offset, offset, i, azimutAngles, horizonAngles)
		horizon.measureHorizonAngle(-offset, distance-offset, numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
		horizon.measureHorizonAngle(-(distance - offset), -offset, 2*numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
		horizon.measureHorizonAngle(offset, -(distance - offset), 3*numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
	}

	return azimutAngles, horizonAngles
}

func (horizon Horizon) computeSkips(distance int) int {
	// Angles changes are the most extreme on the diagonal.
	// Check what size of skips produces acceptable angle changes
	indexDiagonal := int(distance / 2)
	locationDiagonal := horizon.computeOffsetLocation(distance-indexDiagonal, indexDiagonal)
	angleDiagonal := horizon.Place.AzimutAngleTo(locationDiagonal)

	skips := 1
	fineGranularEnough := true
	for ; fineGranularEnough; skips++ {
		locationSkip := horizon.computeOffsetLocation(distance-indexDiagonal-skips, indexDiagonal+skips)
		angleSkip := horizon.Place.AzimutAngleTo(locationSkip)

		fineGranularEnough = math.Abs(angleSkip-angleDiagonal) < 2*math.Pi/float64(horizonAngleResolution)
	}

	return skips - 1
}

func computeSampleResolution(measurementResolution int) int {
	sampleResolution := math.Pow(2, math.Floor(math.Log2(float64(measurementResolution))))
	return int(math.Min(sampleResolution, float64(horizonAngleResolution)))
}

func (horizon Horizon) measureHorizonAngle(stepsNorth, stepsEast, index int, azimutAngles, horizonAngles []float64) {
	sampleLocation := horizon.computeOffsetLocation(stepsNorth, stepsEast)

	azimutAngles[index] = horizon.Place.AzimutAngleTo(sampleLocation)
	horizonAngles[index] = horizon.Place.AltitudeAngleTo(sampleLocation)
}

func (horizon Horizon) computeOffsetLocation(stepsNorth, stepsEast int) *location.Location {
	offsetLocation := location.Location{}
	offsetLocation.Latitude = horizon.Place.Latitude + float64(stepsNorth)*stepSize
	offsetLocation.Longitude = horizon.Place.Longitude + float64(stepsEast)*stepSize

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
