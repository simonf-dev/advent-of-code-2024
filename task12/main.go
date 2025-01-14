package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Position struct {
	positionX int
	positionY int
}

type Fence struct {
	fenceLeft, fenceRight, fenceTop, fenceDown map[Position]struct{}
}

func NewFence() *Fence {
	return &Fence{
		fenceLeft:  make(map[Position]struct{}),
		fenceRight: make(map[Position]struct{}),
		fenceTop:   make(map[Position]struct{}),
		fenceDown:  make(map[Position]struct{}),
	}
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

func sortPositions(data map[Position]struct{}, byX bool) []Position {
	// Extract keys into a slice
	positions := make([]Position, 0, len(data))
	for pos := range data {
		positions = append(positions, pos)
	}

	// Sort the slice using primary and secondary criteria
	sort.Slice(positions, func(i, j int) bool {
		if byX {
			// Primary: positionX; Secondary: positionY
			if positions[i].positionX == positions[j].positionX {
				return positions[i].positionY < positions[j].positionY
			}
			return positions[i].positionX < positions[j].positionX
		}
		// Primary: positionY; Secondary: positionX
		if positions[i].positionY == positions[j].positionY {
			return positions[i].positionX < positions[j].positionX
		}
		return positions[i].positionY < positions[j].positionY
	})

	return positions
}

func countFence(fenceStructure Fence) int {
	fence := 4
	sortedYPositions := sortPositions(fenceStructure.fenceLeft, true)
	indexX, indexY := sortedYPositions[0].positionX, sortedYPositions[0].positionY
	for _, position := range sortedYPositions[1:] {
		if indexX != position.positionX {
			indexX = position.positionX
			indexY = position.positionY
			fence += 1
		} else if indexY+1 != position.positionY {
			fence += 1
			indexY = position.positionY
		} else {
			indexY = position.positionY
		}
	}
	sortedYPositions = sortPositions(fenceStructure.fenceRight, true)
	indexX, indexY = sortedYPositions[0].positionX, sortedYPositions[0].positionY
	for _, position := range sortedYPositions[1:] {
		if indexX != position.positionX {
			indexX = position.positionX
			indexY = position.positionY
			fence += 1
		} else if indexY+1 != position.positionY {
			fence += 1
			indexY = position.positionY
		} else {
			indexY = position.positionY
		}
	}
	sortedXPositions := sortPositions(fenceStructure.fenceTop, false)
	indexX, indexY = sortedXPositions[0].positionX, sortedXPositions[0].positionY
	for _, position := range sortedXPositions[1:] {
		if indexY != position.positionY {
			indexX = position.positionX
			indexY = position.positionY
			fence += 1
		} else if indexX+1 != position.positionX {
			fence += 1
			indexX = position.positionX
		} else {
			indexX = position.positionX
		}
	}
	sortedXPositions = sortPositions(fenceStructure.fenceDown, false)
	indexX, indexY = sortedXPositions[0].positionX, sortedXPositions[0].positionY
	for _, position := range sortedXPositions[1:] {
		if indexY != position.positionY {
			indexX = position.positionX
			indexY = position.positionY
			fence += 1
		} else if indexX+1 != position.positionX {
			fence += 1
			indexX = position.positionX
		} else {
			indexX = position.positionX
		}
	}
	return fence
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

func countArea(fence Fence, visitedMap map[Position]struct{}, lineIndex, colIndex int, field [][]rune) int {
	area := 1
	symbol := field[lineIndex][colIndex]
	visitedMap[Position{positionX: colIndex, positionY: lineIndex}] = struct{}{}
	position := Position{positionX: colIndex, positionY: lineIndex}
	nextPosition := Position{positionX: colIndex, positionY: lineIndex - 1}
	_, contains := visitedMap[nextPosition]
	if lineIndex == 0 || field[lineIndex-1][colIndex] != symbol {
		fence.fenceTop[position] = struct{}{}
	} else if !contains {
		neighborArea := countArea(fence, visitedMap, lineIndex-1, colIndex, field)
		area += neighborArea
	}

	nextPosition = Position{positionX: colIndex - 1, positionY: lineIndex}
	_, contains = visitedMap[nextPosition]
	if colIndex == 0 || field[lineIndex][colIndex-1] != symbol {
		fence.fenceLeft[position] = struct{}{}
	} else if !contains {
		neighborArea := countArea(fence, visitedMap, lineIndex, colIndex-1, field)
		area += neighborArea
	}

	nextPosition = Position{positionX: colIndex, positionY: lineIndex + 1}
	_, contains = visitedMap[nextPosition]
	if lineIndex == len(field)-1 || field[lineIndex+1][colIndex] != symbol {
		fence.fenceDown[Position{positionX: colIndex, positionY: position.positionY + 1}] = struct{}{}
	} else if !contains {
		neighborArea := countArea(fence, visitedMap, lineIndex+1, colIndex, field)
		area += neighborArea
	}

	nextPosition = Position{positionX: colIndex + 1, positionY: lineIndex}
	_, contains = visitedMap[nextPosition]
	if colIndex == len(field[lineIndex])-1 || field[lineIndex][colIndex+1] != symbol {
		fence.fenceRight[Position{positionX: colIndex + 1, positionY: position.positionY}] = struct{}{}
	} else if !contains {
		neighborArea := countArea(fence, visitedMap, lineIndex, colIndex+1, field)
		area += neighborArea
	}

	return area
}

func countCost(field [][]rune) int {
	visitedMap := make(map[Position]struct{})
	result := 0

	for lineIndex, line := range field {
		for colIndex, _ := range line {
			_, contains := visitedMap[Position{positionX: colIndex, positionY: lineIndex}]
			if !contains {
				fence := NewFence()
				area := countArea(*fence, visitedMap, lineIndex, colIndex, field)
				fenceResult := countFence(*fence)
				result += area * fenceResult
			}
		}
	}
	return result
}

func main() {
	field := parseInput("/home/simon/Projects/go-advent-2024/task12/input.txt")
	fmt.Println(countCost(field))
	a()
}
