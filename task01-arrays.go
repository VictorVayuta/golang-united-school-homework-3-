package homework

func average(input [15]float32) (result float32) {
	if len(input) == 0 {
		return float32(0)
	}
	sum := float32(0)
	for _, v := range input {
		sum += v
	}
	return sum / float32(len(input))
}
