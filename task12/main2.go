package main

/*
import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	positionX int
	positionY int
}

func contains(runes []rune, target rune) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(filePath string) [][]rune {
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var arr [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		arr = append(arr, line)
	}
	return arr
}

func countArea(visitedMap map[Position]struct{}, lineIndex, colIndex int, field [][]rune) (int, int) {
	area := 1
	symbol := field[lineIndex][colIndex]
	fence := 0
	visitedMap[Position{positionX: colIndex, positionY: lineIndex}] = struct{}{}
	position := Position{positionX: colIndex, positionY: lineIndex - 1}
	_, contains := visitedMap[position]
	if lineIndex == 0 || field[lineIndex-1][colIndex] != symbol {
		fence += 1
	} else if !contains {
		neighborArea, neighborFence := countArea(visitedMap, lineIndex-1, colIndex, field)
		area += neighborArea
		fence += neighborFence
	}

	position = Position{positionX: colIndex - 1, positionY: lineIndex}
	_, contains = visitedMap[position]
	if colIndex == 0 || field[lineIndex][colIndex-1] != symbol {
		fence += 1
	} else if !contains {
		neighborArea, neighborFence := countArea(visitedMap, lineIndex, colIndex-1, field)
		area += neighborArea
		fence += neighborFence
	}

	position = Position{positionX: colIndex, positionY: lineIndex + 1}
	_, contains = visitedMap[position]
	if lineIndex == len(field)-1 || field[lineIndex+1][colIndex] != symbol {
		fence += 1
	} else if !contains {
		neighborArea, neighborFence := countArea(visitedMap, lineIndex+1, colIndex, field)
		area += neighborArea
		fence += neighborFence
	}

	position = Position{positionX: colIndex + 1, positionY: lineIndex}
	_, contains = visitedMap[position]
	if colIndex == len(field[lineIndex])-1 || field[lineIndex][colIndex+1] != symbol {
		fence += 1
	} else if !contains {
		neighborArea, neighborFence := countArea(visitedMap, lineIndex, colIndex+1, field)
		area += neighborArea
		fence += neighborFence
	}

	return area, fence
}

func countCost(field [][]rune) int {
	visitedMap := make(map[Position]struct{})
	result := 0
	for lineIndex, line := range field {
		for colIndex, _ := range line {
			_, contains := visitedMap[Position{positionX: colIndex, positionY: lineIndex}]
			if !contains {
				area, fence := countArea(visitedMap, lineIndex, colIndex, field)
				result += area * fence
			}
		}
	}
	return result
}

func main() {
	field := parseInput("input.txt")
	fmt.Println(countCost(field))
}
*/
