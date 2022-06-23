package main

import "fmt"

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

func main() {
	averageInput := []float64{1, 2, 3, 4, 5, 6}
	averageRes := average(averageInput)
	fmt.Printf("Average result: %.1f\n", averageRes)

}
