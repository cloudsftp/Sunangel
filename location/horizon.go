package location

import (
	"log"
	"math"
)

const (
	stepSize               float64 = 0.0003
	numSteps               int     = 2 << 11 // 2048
	horizonAngleResolution int     = 1 << 10 // 1024
)

func (loc *Location) ComputeHorizon() {
	log.Printf("Computing horizon for location %f, %f\n", loc.Latitude, loc.Longitude)
	loc.Horizon = [horizonAngleResolution]float64{-math.Pi}

	for k := 1; k <= numSteps; k++ {
		if k%100 == 0 {
			log.Printf("Current sample distance from location: %d\n", k)
		}

		azimutAnglesMeasured, horizonAnglesMeasured := loc.measureHorizonAngles(k)

		azimutAnglesMeasured = azimutAnglesMeasured
		horizonAnglesMeasured = horizonAnglesMeasured

		// TODO: interpolate to max{2 * 2^k, resolution} angles
	}
}

func (loc Location) measureHorizonAngles(distance int) ([]float64, []float64) {
	skips := loc.computeSkips(distance)
	numSamplesPerDiagonal := int(math.Ceil(float64(distance) / float64(skips)))
	numSamples := 4 * numSamplesPerDiagonal

	azimutAngles := make([]float64, numSamples)
	horizonAngles := make([]float64, numSamples)

	for i := 0; i < numSamplesPerDiagonal; i++ {
		offset := i * skips
		loc.measureHorizonAngle(distance-offset, offset, i, azimutAngles, horizonAngles)
		loc.measureHorizonAngle(-offset, distance-offset, numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
		loc.measureHorizonAngle(-(distance - offset), -offset, 2*numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
		loc.measureHorizonAngle(offset, -(distance - offset), 3*numSamplesPerDiagonal+i, azimutAngles, horizonAngles)
	}

	return azimutAngles, horizonAngles
}

func (loc Location) computeSkips(distance int) int {
	// Angles changes are the most extreme on the diagonal.
	// Check what size of skips produces acceptable angle changes
	indexDiagonal := int(distance / 2)
	locationDiagonal := loc.computeOffsetLocation(distance-indexDiagonal, indexDiagonal)
	angleDiagonal := loc.azimutAngleTo(locationDiagonal)

	skips := 1
	fineGranularEnough := true
	for ; fineGranularEnough; skips++ {
		locationSkip := loc.computeOffsetLocation(distance-indexDiagonal-skips, indexDiagonal+skips)
		angleSkip := loc.azimutAngleTo(locationSkip)

		fineGranularEnough = math.Abs(angleSkip-angleDiagonal) < 2*math.Pi/float64(horizonAngleResolution)
	}

	return skips - 1
}

func (loc Location) measureHorizonAngle(stepsNorth, stepsEast, index int, azimutAngles, horizonAngles []float64) {
	sampleLocation := loc.computeOffsetLocation(stepsNorth, stepsEast)

	azimutAngles[index] = loc.azimutAngleTo(sampleLocation)
	horizonAngles[index] = loc.HorizontalAngleTo(sampleLocation)
}

func (loc Location) computeOffsetLocation(stepsNorth, stepsEast int) Location {
	offsetLocation := Location{}
	offsetLocation.Latitude = loc.Latitude + float64(stepsNorth)*stepSize
	offsetLocation.Longitude = loc.Longitude + float64(stepsEast)*stepSize

	return offsetLocation
}
