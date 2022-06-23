package homework

func average(input []float64) float64 {
	if len(input) == 0 {
		return float64(0)
	}
	sum := 0.0
	for _, v := range input {
		sum += v
	}
	return sum / float64(len(input))
}
