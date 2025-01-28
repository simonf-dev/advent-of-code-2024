package main

import (
	"fmt"
	"strings"
)

func ParseStrings(input string, sep string) map[int][]string {
	// Split the input string by ", "
	words := strings.Split(input, sep)

	wordMap := make(map[int][]string)

	for _, word := range words {
		// Trim any leading or trailing whitespace
		word = strings.TrimSpace(word)
		length := len(word)
		wordMap[length] = append(wordMap[length], word)
	}

	return wordMap
}

func CountCombinations(word string, wordMap map[int][]string, startingSize int, cache map[string]int, canBeProducedCache map[string]bool) int {
	if len(word) == 0 {
		return 1
	}
	count := 0
	for i := 1; i <= len(word) && i < startingSize; i++ {
		prefix := word[:i]
		suffix := word[i:]
		countCache, containsSuf := cache[suffix]

		if contains(wordMap[len(prefix)], prefix) && (containsSuf || canBeProduced(suffix, wordMap, startingSize, canBeProducedCache)) {
			if containsSuf {
				count += countCache
			} else {
				countedComb := CountCombinations(suffix, wordMap, startingSize, cache, canBeProducedCache)
				count += countedComb
				cache[suffix] = countedComb
			}
		}
	}

	return count
}

func main() {
	fileContent, _ := readInputFile("input.txt")
	result := ParseStrings(fileContent, ", ")
	printResult(result)
	fileContent, _ = readInputFile("combinations.txt")
	combinations := ParseStrings(fileContent, "\n")
	produceCount := 0
	cache := make(map[string]int)
	canBeProducedCache := make(map[string]bool)
	for _, value := range combinations {
		for _, combination := range value {
			produceCount += CountCombinations(combination, result, len(combination), cache, canBeProducedCache)
		}
	}
	fmt.Println(produceCount)
}
