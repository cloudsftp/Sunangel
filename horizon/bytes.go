package horizon

import (
	"fmt"

	"github.com/cloudsftp/Sunangel/util"
)

// AltitudeToBytes returns the altitude array of a horizon into a byte slice
func (hor *Horizon) AltitudeToBytes() []byte {
	bytes := make([]byte, horizonAngleResolution*util.BytesIn64Bits)

	for i := 0; i < horizonAngleResolution; i++ {
		util.BytesFromFloat64(hor.altitude[i], bytes[i*util.BytesIn64Bits:(i+1)*util.BytesIn64Bits])
	}

	return bytes
}

// AltitudeFromBytes returns an altitude array from a byte slice
func AltitudeFromBytes(bytes []byte) (AltitudeArray, error) {
	a := AltitudeArray{}

	neededBytes := horizonAngleResolution * util.BytesIn64Bits
	if len(bytes) != neededBytes {
		return a, fmt.Errorf("wrong number of bytes need %d, got %d", neededBytes, len(bytes))
	}

	for i := 0; i < horizonAngleResolution; i++ {
		a[i] = util.Float64FromBytes(bytes[i*util.BytesIn64Bits : (i+1)*util.BytesIn64Bits])
	}

	return a, nil
}
