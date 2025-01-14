package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Up    Direction = 0
	Left  Direction = 1
	Down  Direction = 2
	Right Direction = 3
)

type Position struct {
	x, y int
}

func (position Position) Sum(secPosition Position) Position {
	return Position{x: position.x + secPosition.x, y: position.y + secPosition.y}
}

func PrepareShifts() []Position {
	positions := make([]Position, 4)
	positions[3] = Position{x: 1, y: 0}
	positions[2] = Position{x: 0, y: 1}
	positions[1] = Position{x: -1, y: 0}
	positions[0] = Position{x: 0, y: -1}
	return positions
}

type Connection struct {
	Direction    Direction
	StartingNode *Node
	EndingNode   *Node
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	FrontConnections    []Connection
	BackwardConnections []Connection
	NodeDistance        map[Direction]int
	Visited             bool
}

func CountTurningPrice(startingDirection, direction Direction) int {
	if startingDirection == direction {
		return 0
	}
	if (startingDirection+direction)%2 == 1 {
		return 1000
	}
	return 2000
}

func (node *Node) CountMove() {

	for direction, distance := range node.NodeDistance {
		for _, connection := range node.FrontConnections {
			endingDistance, doesExist := connection.EndingNode.NodeDistance[connection.Direction]
			if !doesExist {
				endingDistance = -1
			}
			turningPrice := CountTurningPrice(direction, connection.Direction)
			if endingDistance > turningPrice+1+distance || endingDistance == -1 {
				connection.EndingNode.NodeDistance[connection.Direction] = turningPrice + 1 + distance
			}
		}
	}
}

func (node *Node) CountTiles(direction Direction) {
	fmt.Println("Analyzing node", &node)
	for _, connection := range node.BackwardConnections {
		if connection.Direction == direction {
			for startingDirection, startingDistance := range connection.StartingNode.NodeDistance {
				turningPrice := CountTurningPrice(startingDirection, connection.Direction)
				if turningPrice+1+startingDistance == node.NodeDistance[connection.Direction] {
					connection.StartingNode.Visited = true
					connection.StartingNode.CountTiles(startingDirection)
				}
			}
		}
	}
}

func parseGrid(filePath string) [][]rune {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	return grid
}

func parseGridToNodes(grid [][]rune) (map[Position]*Node, Position, Position) {
	nodeMap := make(map[Position]*Node)
	var endPosition Position
	var startPosition Position
	for lineIndex, line := range grid {
		for colIndex, char := range line {
			switch char {
			case '#':
				continue
			case '.':
				nodeMap[Position{x: colIndex, y: lineIndex}] = &Node{FrontConnections: make([]Connection, 0), BackwardConnections: make([]Connection, 0), NodeDistance: make(map[Direction]int), Visited: false}
			case 'S':
				distances := make(map[Direction]int)
				distances[Right] = 0
				nodeMap[Position{x: colIndex, y: lineIndex}] = &Node{FrontConnections: make([]Connection, 0), BackwardConnections: make([]Connection, 0), NodeDistance: distances, Visited: true}
				startPosition = Position{x: colIndex, y: lineIndex}

			case 'E':
				nodeMap[Position{x: colIndex, y: lineIndex}] = &Node{FrontConnections: make([]Connection, 0), BackwardConnections: make([]Connection, 0), NodeDistance: make(map[Direction]int), Visited: true}
				endPosition = Position{x: colIndex, y: lineIndex}
			}
		}
	}
	preparedShifts := PrepareShifts()
	for value := range nodeMap {
		for shiftDirection, shift := range preparedShifts {
			newPosition := shift.Sum(value)
			otherNode, contains := nodeMap[newPosition]
			if contains {
				currentNode := nodeMap[value]
				connection := Connection{
					StartingNode: currentNode,
					EndingNode:   otherNode,
					Direction:    Direction(shiftDirection),
				}
				currentNode.FrontConnections = append(currentNode.FrontConnections, connection)
				otherNode.BackwardConnections = append(otherNode.BackwardConnections, connection)
				nodeMap[value] = currentNode
			}
		}
	}
	return nodeMap, startPosition, endPosition
}

func Min(list map[Direction]int) int {
	if len(list) == 0 {
		return -1
	}

	min := -1
	for _, value := range list {
		if value < min || min == -1 {
			min = value
		}
	}

	return min
}

func Sum(list map[Direction]int) int {
	if len(list) == 0 {
		return 0
	}

	sum := 0
	for _, value := range list {
		sum += value
	}

	return sum
}
func PrintMap(fieldMap map[Position]*Node, tall, width int) {
	for y := range tall {
		for x := range width {
			position := Position{x: x, y: y}
			value, contains := fieldMap[position]
			if contains {
				fmt.Printf("%-5d", value.Visited)
			} else {
				fmt.Print("#####")
			}
		}
		fmt.Print("\n")
	}
}
func main() {
	gridRunes := parseGrid("input.txt")
	gridNodes, _, endPosition := parseGridToNodes(gridRunes)
	for range 2000 {
		for _, value := range gridNodes {
			value.CountMove()
		}
	}
	minDistance := Min(gridNodes[endPosition].NodeDistance)
	PrintMap(gridNodes, 17, 17)
	for key, value := range gridNodes[endPosition].NodeDistance {
		if value == minDistance {
			gridNodes[endPosition].CountTiles(key)
		}
	}
	visited := 0
	for _, value := range gridNodes {
		if value.Visited {
			visited += 1
		}
	}
	fmt.Println(visited)
	PrintMap(gridNodes, 17, 17)
}
