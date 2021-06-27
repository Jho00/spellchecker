package internal

import (
	"fmt"
	"strings"
)

func isSimilarWord_PrimitiveWay(checkableWord string, testWord string) bool {
	return strings.ToLower(checkableWord) == strings.ToLower(testWord)
}

func chunkingArray(inputArray []string, chunkSize int) [][]string  {
	var divided [][]string

	for i := 0; i < len(inputArray); i += chunkSize {
		end := i + chunkSize

		if end > len(inputArray) {
			end = len(inputArray)
		}

		divided = append(divided, inputArray[i:end])
	}

	return divided
}

func unionMaps(base map[int][]int, target map[int][]int) map[int][]int  {
	for key, value := range target {
		_, ok := base[key]

		if ok {
			fmt.Println("WARNING: Overwrite key", key, "with values: ", value)
		}

		base[key] = value
	}


	return base
}