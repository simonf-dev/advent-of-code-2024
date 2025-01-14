package main

import (
	"bufio"
	"fmt"
	"os"
)

type Field int

const (
	Empty Field = iota
	Robot
	LeftBox
	RightBox
	Barrier
)

type Direction int

const (
	Right Direction = iota
	Left
	Top
	Down
)

func decideMove(direction Direction) (int, int) {
	switch direction {
	case Top:
		return -1, 0
	case Down:
		return 1, 0
	case Right:
		return 0, 1
	case Left:
		return 0, -1
	default:
		panic("Invalid direction")
	}
}
func CheckIfCanMove(direction Direction, x, y int, fields [][]Field) bool {
	if x <= 0 || y <= 0 || y >= len(fields) || x >= len(fields[0]) {
		panic("Trying to move from bad field")
	}
	newY, newX := decideMove(direction)
	newY += y
	newX += x
	if newX <= 0 || newY <= 0 || newY >= len(fields) || newX >= len(fields[0]) {
		return false
	}
	fieldValue := fields[newY][newX]
	switch fieldValue {
	case Barrier:
		return false
	case Empty:
		return true
	case Robot:
		panic("Cannot move object to robot, something is wrong")
	case LeftBox:
		switch direction {
		case Right, Left:
			return CheckIfCanMove(direction, newX, newY, fields)
		case Top, Down:
			return CheckIfCanMove(direction, newX+1, newY, fields) && CheckIfCanMove(direction, newX, newY, fields)
		}
	case RightBox:
		switch direction {
		case Right, Left:
			return CheckIfCanMove(direction, newX, newY, fields)
		case Top, Down:
			return CheckIfCanMove(direction, newX-1, newY, fields) && CheckIfCanMove(direction, newX, newY, fields)
		}

	}
	panic("Invalid combination of params")
}

func moveField(direction Direction, x, y int, fields [][]Field) (int, int) {
	if x <= 0 || y <= 0 || y >= len(fields) || x >= len(fields[0]) {
		panic("Trying to move from bad field")
	}
	newY, newX := decideMove(direction)
	newY += y
	newX += x
	if newX <= 0 || newY <= 0 || newY >= len(fields) || newX >= len(fields[0]) {
		return x, y
	}
	fieldValue := fields[newY][newX]
	switch fieldValue {
	case Barrier:
		return x, y
	case Empty:
		fields[newY][newX] = fields[y][x]
		fields[y][x] = Empty
		return newX, newY
	case Robot:
		panic("Cannot move object to robot, something is wrong")
	case LeftBox:
		switch direction {
		case Right, Left:
			moveField(direction, newX, newY, fields)
			if fields[newY][newX] == Empty {
				fields[newY][newX] = fields[y][x]
				fields[y][x] = Empty
				return newX, newY
			}
		case Top, Down:
			if CheckIfCanMove(direction, newX+1, y, fields) && CheckIfCanMove(direction, newX, newY, fields) {
				moveField(direction, newX, newY, fields)
				moveField(direction, newX+1, newY, fields)
				if fields[newY][newX] == Empty {
					fields[newY][newX] = fields[y][x]
					fields[y][x] = Empty
					return newX, newY
				}
			}
		}
	case RightBox:
		switch direction {
		case Right, Left:
			moveField(direction, newX, newY, fields)
			if fields[newY][newX] == Empty {
				fields[newY][newX] = fields[y][x]
				fields[y][x] = Empty
				return newX, newY
			}
		case Top, Down:
			if CheckIfCanMove(direction, newX-1, newY, fields) && CheckIfCanMove(direction, newX, newY, fields) {
				moveField(direction, newX, newY, fields)
				moveField(direction, newX-1, newY, fields)
				if fields[newY][newX] == Empty {
					fields[newY][newX] = fields[y][x]
					fields[y][x] = Empty
					return newX, newY
				}
			}
		}
	}
	return x, y
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseField(filePath string) ([][]Field, int, int) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	field := make([][]Field, 0)
	robotX, robotY := 0, 0
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		line := make([]Field, len(text)*2)
		x := 0
		for _, char := range text {
			switch char {
			case '#':
				line[x] = Barrier
				line[x+1] = Barrier
			case '@':
				line[x] = Robot
				line[x+1] = Empty
				robotX, robotY = x, y
			case '.':
				line[x] = Empty
				line[x+1] = Empty
			case 'O':
				line[x] = LeftBox
				line[x+1] = RightBox
			}
			x += 2
		}
		y += 1
		field = append(field, line)
	}
	return field, robotX, robotY
}

func ParseMoves(filePath string) []Direction {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	directions := make([]Direction, 0)
	for scanner.Scan() {
		text := scanner.Text()
		for _, char := range text {
			switch char {
			case '>':
				directions = append(directions, Right)
			case '^':
				directions = append(directions, Top)
			case '<':
				directions = append(directions, Left)
			case 'v':
				directions = append(directions, Down)
			}
		}
	}
	return directions
}

func CountResult(field [][]Field) int {
	result := 0
	for y := range field {
		for x := range field[y] {
			if field[y][x] == LeftBox {
				result += y*100 + x
			}
		}
	}
	return result
}

func PrintField(field [][]Field) {
	for y := range field {
		for x := range field[y] {
			switch field[y][x] {
			case LeftBox:
				fmt.Print("[")
			case Barrier:
				fmt.Print("#")
			case RightBox:
				fmt.Print("]")
			case Robot:
				fmt.Print("@")
			case Empty:
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}
func main() {
	field, robotX, robotY := ParseField("plan.txt")
	directions := ParseMoves("moves.txt")
	for _, direction := range directions {
		robotX, robotY = moveField(direction, robotX, robotY, field)
		//PrintField(field)
	}
	fmt.Println(CountResult(field))

}
