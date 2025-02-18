package homework

import (
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
