package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type location struct {
	positionX int
	positionY int
}

type fieldInfo struct {
	antennas  map[rune][]location
	antinodes map[location]struct{}
	sizeX     int
	sizeY     int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseField(filePath string) fieldInfo {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	field := fieldInfo{antinodes: make(map[location]struct{}), antennas: make(map[rune][]location)}
	lineIndex := 0
	for scanner.Scan() {
		strippedLine := strings.TrimSpace(scanner.Text())
		for rowIndex, char := range strippedLine {
			field.sizeX = rowIndex + 1
			if char == '.' {
				continue
			}
			_, contains := field.antennas[char]
			if contains {
				field.antennas[char] = append(field.antennas[char], location{positionX: rowIndex, positionY: lineIndex})
			} else {
				field.antennas[char] = []location{location{positionX: rowIndex, positionY: lineIndex}}
			}
		}
		lineIndex += 1
	}
	field.sizeY = lineIndex
	return field
}

func countAntinnode(firstAntenna, secondAntenna location, positionList *map[location]struct{}, sizeX, sizeY int) {
	differenceX, differenceY := firstAntenna.positionX-secondAntenna.positionX, firstAntenna.positionY-secondAntenna.positionY
	firstPosition := location{positionX: firstAntenna.positionX, positionY: firstAntenna.positionY}
	for sizeY > firstPosition.positionY && firstPosition.positionY >= 0 && sizeX > firstPosition.positionX && firstPosition.positionX >= 0 {
		(*positionList)[firstPosition] = struct{}{}
		firstPosition = location{positionX: firstPosition.positionX + differenceX, positionY: firstPosition.positionY + differenceY}
	}

	secondPosition := location{positionX: secondAntenna.positionX, positionY: secondAntenna.positionY}
	for sizeY > secondPosition.positionY && secondPosition.positionY >= 0 && sizeX > secondPosition.positionX && secondPosition.positionX >= 0 {
		(*positionList)[secondPosition] = struct{}{}
		secondPosition = location{positionX: secondPosition.positionX - differenceX, positionY: secondPosition.positionY - differenceY}
	}
}

func countAntinnodes(field *fieldInfo) {
	for _, antennaLocations := range field.antennas {
		for index, position := range antennaLocations {
			for combinationIndex := index + 1; combinationIndex < len(antennaLocations); combinationIndex++ {
				countAntinnode(position, antennaLocations[combinationIndex], &field.antinodes, field.sizeX, field.sizeY)
			}
		}
	}
}

func printField(field fieldInfo) {
	for rowIndex := range field.sizeY {
		for columnIndex := range field.sizeX {
			_, contains := field.antinodes[location{positionX: columnIndex, positionY: rowIndex}]
			if contains {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
func main() {
	output := parseField("input.txt")
	countAntinnodes(&output)
	fmt.Println(len(output.antinodes))
	printField(output)
}
