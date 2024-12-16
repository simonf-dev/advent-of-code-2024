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

var xmasRunes []rune = []rune("XMAS")

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

func checkRow(lineIndex int, colIndex int, structure [][]rune) int {
	foundCount := 2
	for i := range 4 {
		if colIndex+i >= len(structure[i]) || structure[lineIndex][colIndex+i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	for i := range 4 {
		if colIndex-i < 0 || structure[lineIndex][colIndex-i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	return foundCount
}

func checkColumn(lineIndex int, colIndex int, structure [][]rune) int {
	foundCount := 2
	for i := range 4 {
		if lineIndex+i >= len(structure) || structure[lineIndex+i][colIndex] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	for i := range 4 {
		if lineIndex-i < 0 || structure[lineIndex-i][colIndex] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	return foundCount
}

func checkDiagonals(lineIndex int, colIndex int, structure [][]rune) int {
	foundCount := 4
	for i := range 4 {
		if lineIndex+i >= len(structure) || colIndex+i >= len(structure[lineIndex+i]) {
			foundCount -= 1
			break
		}
		if structure[lineIndex+i][colIndex+i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	for i := range 4 {
		if lineIndex+i >= len(structure) || colIndex-i < 0 {
			foundCount -= 1
			break
		}
		if structure[lineIndex+i][colIndex-i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	for i := range 4 {
		if lineIndex-i < 0 || colIndex-i < 0 {
			foundCount -= 1
			break
		}
		if  structure[lineIndex-i][colIndex-i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}
	for i := range 4 {
		if lineIndex-i < 0 || colIndex+i >= len(structure[lineIndex-i]) {
			foundCount -= 1
			break
		}
		if lineIndex-i < 0 || structure[lineIndex-i][colIndex+i] != xmasRunes[i] {
			foundCount -= 1
			break
		}
	}

	return foundCount
}

func mainIteration(structure [][]rune) int {
	foundCount := 0
	for lineIndex, _ := range structure {
		for colIndex, _ := range structure[lineIndex] {
			foundCount += checkRow(lineIndex, colIndex, structure)
			foundCount += checkColumn(lineIndex, colIndex, structure)
			foundCount += checkDiagonals(lineIndex, colIndex, structure)
		}
	}
	return foundCount
}

func main() {
	parsedList := parseInput("input.txt")
	fmt.Println(mainIteration(parsedList))
}
