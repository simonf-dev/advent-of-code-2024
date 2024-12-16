package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// BubbleSort sorts a slice of integers in ascending order using the Bubble Sort algorithm.
func BubbleSort(arr []int, rules map[int][]int) {
	// Get the length of the slice
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			value, contains := rules[arr[j+1]]
			if !contains {
				continue
			}
			if hasOverlap([]int{arr[j]}, value) {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func hasOverlap(slice1, slice2 []int) bool {
	elementMap := make(map[int]bool)

	for _, val := range slice1 {
		elementMap[val] = true
	}

	for _, val := range slice2 {
		if _, exists := elementMap[val]; exists {
			return true // Overlap found
		}
	}

	return false // No overlap
}
func parseUpdates(filePath string) [][]int {
	dat, err := os.Open(filePath)
	check(err)
	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)
	var updateList [][]int
	for scanner.Scan() {
		updateSlice := strings.Split(scanner.Text(), ",")
		var updateLine []int
		for _, update := range updateSlice {
			updateValue, err := strconv.Atoi(update)
			check(err)
			updateLine = append(updateLine, updateValue)
		}
		updateList = append(updateList, updateLine)
	}
	return updateList
}

func parseRules(filePath string) map[int][]int {
	dat, err := os.Open(filePath)
	check(err)
	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)
	rulesMap := make(map[int][]int)
	for scanner.Scan() {
		ruleSlice := strings.Split(scanner.Text(), "|")
		if len(ruleSlice) != 2 {
			panic(fmt.Sprintf("ruleSlice should be 2, but it is %d", len(ruleSlice)))
		}
		ruleValueInt, err := strconv.Atoi(ruleSlice[0])
		check(err)
		valueInt, err := strconv.Atoi(ruleSlice[1])
		check(err)
		bufferSlice, contains := rulesMap[ruleValueInt]
		if contains {
			rulesMap[ruleValueInt] = append(bufferSlice, valueInt)
		} else {
			rulesMap[ruleValueInt] = []int{valueInt}
		}
	}
	return rulesMap
}

func checkRule(rules map[int][]int, updateLine []int) int {
	updatesExecuted := []int{}
	for _, update := range updateLine {
		updateRules, contains := rules[update]
		if contains && hasOverlap(updatesExecuted, updateRules) {
			BubbleSort(updateLine, rules)
			return updateLine[len(updateLine)/2]
		}
		updatesExecuted = append(updatesExecuted, update)
	}
	return updateLine[len(updateLine)/2]
}
func checkRules(rules map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		result += checkRule(rules, update)
	}
	return result
}
func main() {
	parsedRules := parseRules("rules.txt")
	parsedUpdates := parseUpdates("updates.txt")
	fmt.Println(checkRules(parsedRules, parsedUpdates))
}
