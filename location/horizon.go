package location

import (
	"log"
	"math"

	"github.com/cloudsftp/Sunangel/util"
)

const (
	stepSize               float64 = 0.0003
	numSteps               int     = 2 << 11 // 2048
	horizonAngleResolution int     = 1 << 10 // 1024
	horizonAngleWidth      float64 = 2 * math.Pi / float64(horizonAngleResolution)
)

func (loc *Location) computeHorizon() {
	log.Printf("Computing horizon for location %f, %f\n", loc.Latitude, loc.Longitude)
	loc.Horizon = [horizonAngleResolution]float64{}
	for i := 0; i < len(loc.Horizon); i++ {
		loc.Horizon[i] = -math.Pi
	}

	for k := 1; k <= numSteps; k++ {
		azimutAnglesMeasured, horizonAnglesMeasured := loc.measureHorizonAngles(k)

		currHorizonAngleResolution := int(math.Min(math.Pow(2, float64(k)), float64(horizonAngleResolution)))
		loc.interpolateHorizonAnglesFromSamples(azimutAnglesMeasured, horizonAnglesMeasured, currHorizonAngleResolution)
		if currHorizonAngleResolution < horizonAngleResolution {
			loc.interpolateHorizonAnglesFromHorizonAngles(currHorizonAngleResolution)
		}
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

func (loc *Location) interpolateHorizonAnglesFromSamples(azimutAngles, horizonAngles []float64, resolution int) {
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
		if horizonAngle > loc.Horizon[i*skips] {
			loc.Horizon[i*skips] = horizonAngle
		}
	}
}

func (loc *Location) interpolateHorizonAnglesFromHorizonAngles(resolution int) {
	skipsInterpolated := horizonAngleResolution / resolution // resolution always power of 2 and < horizonAngleResolution
	offsetNotInterpolated := skipsInterpolated / 2

	for i := 0; i < resolution; i++ {
		leftIndex := i * skipsInterpolated
		rightIndex := (leftIndex + skipsInterpolated) % len(loc.Horizon)
		tgtIndex := leftIndex + offsetNotInterpolated

		leftHorizonAngle := loc.Horizon[leftIndex]
		rightHorizonAngle := loc.Horizon[rightIndex]

		horizonAngle := (rightHorizonAngle - leftHorizonAngle) / 2
		horizonAngle += leftHorizonAngle

		if horizonAngle > loc.Horizon[tgtIndex] {
			loc.Horizon[tgtIndex] = horizonAngle
		}
	}
}

func (loc Location) GetHorizonAngleAt(tgtAzimutAngle float64) float64 {
	leftIndex := int(tgtAzimutAngle / horizonAngleWidth)
	rightIndex := (leftIndex + 1) % horizonAngleResolution

	leftAzimutAngle := float64(leftIndex) * horizonAngleWidth
	rightAzimutAngle := float64(rightIndex) * horizonAngleWidth

	leftHorizonAngle := loc.Horizon[leftIndex]
	rightHorizonAngle := loc.Horizon[rightIndex]

	return util.LinInt(tgtAzimutAngle, leftAzimutAngle, leftHorizonAngle, rightAzimutAngle, rightHorizonAngle)
}
