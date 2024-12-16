package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var xmasRunes []rune = []rune("MAS")

func parseInput(filePath string) [][]rune {
	file, err := os.Open(filePath)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var arr [][]rune
	for scanner.Scan() {
		arr = append(arr, []rune(scanner.Text()))
	}
	file.Close()
	return arr
}

func checkCross(lineIndex int, colIndex int, structure [][] rune) int {
	if lineIndex <= 0 || lineIndex >= len(structure) - 1 || colIndex <= 0 || colIndex >= len(structure[lineIndex]) - 1 || structure[lineIndex][colIndex] != xmasRunes[1] {
		return 0
	}
	if (structure[lineIndex-1][colIndex-1] != xmasRunes[0] || structure[lineIndex+1][colIndex+1] != xmasRunes[2]) && (structure[lineIndex-1][colIndex-1] != xmasRunes[2] || structure[lineIndex+1][colIndex+1] != xmasRunes[0]) {
		return 0
	}
        if (structure[lineIndex-1][colIndex+1] != xmasRunes[0] || structure[lineIndex+1][colIndex-1] != xmasRunes[2]) && (structure[lineIndex-1][colIndex+1] != xmasRunes[2] || structure[lineIndex+1][colIndex-1] != xmasRunes[0]) {
                return 0
        }
	return 1
	
}

func mainIteration(structure [][]rune) int {
	foundCount := 0
	for lineIndex, _ := range structure {
		for colIndex, _ := range structure[lineIndex] {
			foundCount += checkCross(lineIndex, colIndex, structure)
		}
	}
	return foundCount
}

func main() {
	parsedList := parseInput("input.txt")
	fmt.Println(mainIteration(parsedList))
}
