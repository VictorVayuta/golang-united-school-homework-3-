package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input *map[int]string) []string {
	var output []string
	var keys []int
	for k := range *input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		output = append(output, (*input)[k])
	}
	return output
}

func main() {
	sortMapValuesInput := map[int]string{2: "a", 0: "b", 1: "c"}
	sortMapValuesRes := sortMapValues(&sortMapValuesInput)
	fmt.Printf("sortMapValues result: %v\n", sortMapValuesRes)

}
