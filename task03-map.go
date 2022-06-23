package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	var output []string
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		output = append(output, (input)[k])
	}
	return output
}

func main() {

	sortMapValuesInput := map[int]string{2: "a", 0: "b", 1: "c"}
	sortMapValuesRes := sortMapValues(sortMapValuesInput)
	sortMapValuesInput1 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	sortMapValuesRes1 := sortMapValues(sortMapValuesInput1)
	fmt.Printf("Sort map values first result: %v\n", sortMapValuesRes)
	fmt.Printf("Sort map values second result: %v\n", sortMapValuesRes1)

}
