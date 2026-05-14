package features



func computeIATStats(iats []float64) (mean float64, variance float64) {

	n := len(iats)
	if n == 0 {
		return 0, 0
	}

	sum := 0.0
	for _, v := range iats {
		sum += v
	}
	mean = sum / float64(n)

	varSum := 0.0
	for _, v := range iats {
		diff := v - mean
		varSum += diff * diff
	}

	variance = varSum / float64(n)
	return
}

