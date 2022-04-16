package util

// LinInt returns a linear interpolation of two points
func LinInt(tgtX, lX, lY, rX, rY float64) float64 {
	result := tgtX - lX
	result *= rY - lY
	result /= rX - lX
	result += lY

	return result
}
