package internal

import (
	"fmt"
	"spellhecker/pkg"
	"strings"
)

func Check(text string, config pkg.Config) {
	result := make(map[int][]int)
	var partialResult map[int][]int

	splitted := strings.Split(text, " ")

	partialResult = launchTautologiesSearch(
		chunkingArray(splitted, config.TautologyCheckDistance),
		0)

	result = unionMaps(result, partialResult)

	partialResult = launchTautologiesSearch(
		chunkingArray(splitted[config.TautologyCheckOffset:], config.TautologyCheckDistance),
		config.TautologyCheckOffset)

	result = unionMaps(result, partialResult)

	for key, value := range result {
		fmt.Println("For key", key, "found tautologies: ", value)
	}
}

func launchTautologiesSearch(chunks [][]string, offset int) map[int][]int {
	result := make(map[int][]int)

	chunkSize := len(chunks[0])

	for i, ch := range chunks {
		found := findTautologies(ch, i * chunkSize + offset)

		for key, element := range found {
			result[key] = element
		}
	}

	return result
}

func findTautologies(words []string, offset int) map[int][]int {
	tautologies := make(map[int][]int)

	getAbsoluteValue := func(value int) int { return value + offset}

	for i, word := range words {
		for j, secondRoundWord := range words {
			if i == j {
				continue
			}

			if !isSimilarWord_PrimitiveWay(word, secondRoundWord) {
				continue
			}

			_, ok := tautologies[getAbsoluteValue(j)]

			if !ok {
				arr, ok := tautologies[getAbsoluteValue(i)]

				if ok {
					tautologies[getAbsoluteValue(i)] = append(arr, getAbsoluteValue(j))
				} else {
					tautologies[getAbsoluteValue(i)] = []int{getAbsoluteValue(j)}
				}
			}
		}
	}

	return purifyFoundTautologies(tautologies)
}

func purifyFoundTautologies(found map[int][]int) map[int][]int {
	for _, element := range found {
		for _, value := range element {
			_, ok := found[value]

			if ok {
				delete(found, value)
			}
		}
	}

	return found
}
