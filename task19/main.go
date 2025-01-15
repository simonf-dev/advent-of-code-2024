package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileContent, _ := readInputFile("input.txt")
	result := processStrings(fileContent, ", ")
	printResult(result)
	fileContent, _ = readInputFile("combinations.txt")
	combinations := processStrings(fileContent, "\n")
	produceCount := 0
	for _, value := range combinations {
		for _, combination := range value {
			if canBeProduced(combination, result, len(combination)) {
				produceCount += 1
			}
		}
	}
	fmt.Println(produceCount)
}

func readInputFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}
		content.Write(buffer[:bytesRead])
		if err == io.EOF {
			break
		}
	}
	return content.String(), nil
}

func processStrings(input string, sep string) map[int][]string {
	// Split the input string by ", "
	words := strings.Split(input, sep)

	wordMap := make(map[int][]string)

	for _, word := range words {
		// Trim any leading or trailing whitespace
		word = strings.TrimSpace(word)
		length := len(word)
		wordMap[length] = append(wordMap[length], word)
	}

	for length := 2; length <= len(words); length++ {
		wordMap[length] = filterWords(wordMap[length], wordMap)
		if len(wordMap[length]) == 0 {
			delete(wordMap, length)
		}
	}

	return wordMap
}

func filterWords(words []string, wordMap map[int][]string) []string {
	var result []string
	for _, word := range words {
		if !canBeProduced(word, wordMap, len(word)) {
			result = append(result, word)
		}
	}
	return result
}

func canBeProduced(word string, wordMap map[int][]string, startingSize int) bool {
	if len(word) == 0 {
		return true
	}

	for i := 1; i <= len(word) && i < startingSize; i++ {
		prefix := word[:i]
		suffix := word[i:]

		if contains(wordMap[len(prefix)], prefix) && canBeProduced(suffix, wordMap, startingSize) {
			return true
		}
	}

	return false
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func printResult(result map[int][]string) {
	for length, words := range result {
		fmt.Printf("Length %d: %v\n", length, words)
	}
}
