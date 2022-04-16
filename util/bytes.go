package util

import (
	"encoding/binary"
	"math"
)

const BytesIn64Bits = 8

func Float64FromBytes(bytes []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytes))
}

func BytesFromFloat64(val float64, bytes []byte) {
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(val))
}
