package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stoneCache struct {
	depth       int
	stoneNumber int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func parseStones(filePath string, blinks int) int {
	cacheMap := make(map[stoneCache]int)
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	stonesCount := 0
	for scanner.Scan() {
		stoneNumber, err := strconv.Atoi(scanner.Text())
		check(err)
		fmt.Println(countStone(stoneNumber, blinks, &cacheMap))
	}
	return stonesCount
}

func countStone(stoneNumber, depth int, cacheMap *map[stoneCache]int) int {
	if depth == 0 {
		return 1
	}
	value, contains := (*cacheMap)[stoneCache{depth: depth, stoneNumber: stoneNumber}]
	if contains {
		return value
	}
	numberStr := strconv.Itoa(stoneNumber)
	returnValue := 0
	if numberStr == "0" {
		returnValue = countStone(1, depth-1, cacheMap)
	} else if len(numberStr)%2 == 0 {
		mid := len(numberStr) / 2
		firstHalfStr := numberStr[:mid]
		secondHalfStr := numberStr[mid:]

		firstHalf, err := strconv.Atoi(firstHalfStr)
		check(err)
		secondHalf, err := strconv.Atoi(secondHalfStr)
		check(err)
		returnValue += countStone(firstHalf, depth-1, cacheMap)
		returnValue += countStone(secondHalf, depth-1, cacheMap)

	} else {
		returnValue += countStone(stoneNumber*2024, depth-1, cacheMap)
	}
	(*cacheMap)[stoneCache{depth: depth, stoneNumber: stoneNumber}] = returnValue
	return returnValue
}

func main() {
	fmt.Println(parseStones("input.txt", 100))

}
