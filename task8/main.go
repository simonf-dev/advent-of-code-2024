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
			if char == '.' {
				continue
			}
			_, contains := field.antennas[char]
			if contains {
				field.antennas[char] = append(field.antennas[char], location{positionX: rowIndex, positionY: lineIndex})
			} else {
				field.antennas[char] = []location{location{positionX: rowIndex, positionY: lineIndex}}
			}
			field.sizeX = rowIndex + 1
		}
	}
	field.sizeY = lineIndex + 1
	return field
}

func countAntinnode(firstAntenna, secondAntenna location, positionList *map[location]struct{}, sizeX, sizeY int) {
	differenceX, differenceY := firstAntenna.positionX-secondAntenna.positionX, firstAntenna.positionY-secondAntenna.positionY
	firstPosition := location{positionX: firstAntenna.positionX + differenceX, positionY: firstAntenna.positionY + differenceY}
	secondPosition := location{positionX: secondAntenna.positionX - differenceX, positionY: secondAntenna.positionY - differenceY}
	if sizeY > firstPosition.positionY && firstPosition.positionY >= 0 && sizeX > firstPosition.positionX && firstPosition.positionX >= 0 {
		(*positionList)[firstPosition] = struct{}{}
	}
	if sizeY > secondPosition.positionY && secondPosition.positionY >= 0 && sizeX > secondPosition.positionX && secondPosition.positionX >= 0 {
		(*positionList)[firstPosition] = struct{}{}
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
func main() {
	output := parseField("input.txt")
	fmt.Println(output)
	countAntinnodes(&output)
	fmt.Println(len(output.antinodes))

}
