package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type trailPoint struct {
	height          int
	connectedPoints []*trailPoint
}

func (point *trailPoint) checkTrail() int {
	if point.height == 9 {
		return 1
	}
	sumCount := 0
	for _, value := range point.connectedPoints {
		if value.height == point.height+1 {
			sumCount += value.checkTrail()
		}
	}
	return sumCount
}

func (point *trailPoint) countPeak(peakMap *map[*trailPoint]struct{}) {
	if point.height == 9 {
		(*peakMap)[point] = struct{}{}
		return
	}
	for _, value := range point.connectedPoints {
		if value.height == point.height+1 {
			value.countPeak(peakMap)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connectPoint(point *trailPoint, trailPoints [][]trailPoint, lineIndex, rowIndex int) {
	connectedPoints := make([]*trailPoint, 0)
	if lineIndex > 0 && trailPoints[lineIndex-1][rowIndex].height == point.height+1 {
		connectedPoints = append(connectedPoints, &trailPoints[lineIndex-1][rowIndex])
	}
	if rowIndex > 0 && trailPoints[lineIndex][rowIndex-1].height == point.height+1 {
		connectedPoints = append(connectedPoints, &trailPoints[lineIndex][rowIndex-1])
	}
	if lineIndex < len(trailPoints)-1 && trailPoints[lineIndex+1][rowIndex].height == point.height+1 {
		connectedPoints = append(connectedPoints, &trailPoints[lineIndex+1][rowIndex])
	}
	if rowIndex < len(trailPoints[lineIndex])-1 && trailPoints[lineIndex][rowIndex+1].height == point.height+1 {
		connectedPoints = append(connectedPoints, &trailPoints[lineIndex][rowIndex+1])
	}
	point.connectedPoints = connectedPoints
}
func connectPoints(trailPoints [][]trailPoint) []*trailPoint {
	filteredTrailPoints := make([]*trailPoint, 0)
	for lineIndex, line := range trailPoints {
		for rowIndex, _ := range line {
			connectPoint(&trailPoints[lineIndex][rowIndex], trailPoints, lineIndex, rowIndex)
			if trailPoints[lineIndex][rowIndex].height == 0 {
				filteredTrailPoints = append(filteredTrailPoints, &trailPoints[lineIndex][rowIndex])
			}
		}
	}
	return filteredTrailPoints
}
func parseTrailMap(filePath string) []*trailPoint {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	trailPoints := make([][]trailPoint, 0)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		trailPointsLine := make([]trailPoint, 0)
		for _, r := range text {
			height, err := strconv.Atoi(string(r))
			check(err)
			trailPointsLine = append(trailPointsLine, trailPoint{height: height})
		}
		trailPoints = append(trailPoints, trailPointsLine)
	}
	filteredTrailPoints := connectPoints(trailPoints)
	return filteredTrailPoints
}

func main() {
	filteredStarts := parseTrailMap("input.txt")
	var sumTrails = 0
	for _, value := range filteredStarts {
		sumTrails += value.checkTrail()
	}
	fmt.Println(sumTrails)
}
