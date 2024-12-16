
import (
	"bufio"
	"fmt"
	"os"
)

type Direction rune

const (
	Up    Direction = '^'
	Down  Direction = 'ˇ'
	Left  Direction = '<'
	Right Direction = '>'
)

type Position struct {
	positionX int
	positionY int
}
type Guard struct {
	positionX int
	positionY int
	direction Direction
}

const barrier rune = '#'
const empty rune = '.'
const visitedField rune = 'X'

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

func parseInput(filePath string) ([][]rune, Guard) {
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	symbols := []rune{'^', '<', '>', 'ˇ'}
	var arr [][]rune
	var guard *Guard
	for scanner.Scan() {
		line := []rune(scanner.Text())
		arr = append(arr, line)
		for index, positionValue := range line {
			if contains(symbols, positionValue) {
				guard = &Guard{positionX: index, positionY: len(arr) - 1, direction: Direction(positionValue)}
			}
		}
	}
	if guard == nil {
		panic("Guard wasn't initialized during input parsing")
	}
	return arr, *guard
}

func countVisited(fields [][]rune) int {
	visited := 0
	for _, line := range fields {
		for _, field := range line {
			switch field {
			case '<', '>', 'ˇ', '^', visitedField:
				visited += 1
			}
		}
	}
	return visited
}

func isOutside(fields [][]rune, positionX, positionY int) bool {
	if positionX < 0 || positionY < 0 || positionY >= len(fields) || positionX >= len(fields[positionY]) {
		return true
	}
	return false
}

func isBarrier(fields [][]rune, positionX, positionY int) bool {
	if isOutside(fields, positionX, positionY) || fields[positionY][positionX] != barrier {
		return false
	}
	return true
}

func decideDirection(direction Direction) Direction {
	switch direction {
	case Direction('>'):
		return Direction('ˇ')
	case Direction('ˇ'):
		return Direction('<')
	case Direction('<'):
		return Direction('^')
	case Direction('^'):
		return Direction('>')
	}
	panic("Not known direction")
}
func moveGuard(fields [][]rune, guard *Guard) {
	orPositionX, orPositionY := (*guard).positionX, (*guard).positionY
	switch guard.direction {
	case Direction('>'):
		(*guard).positionX += 1
	case Direction('<'):
		(*guard).positionX -= 1
	case Direction('^'):
		(*guard).positionY -= 1
	case Direction('ˇ'):
		(*guard).positionY += 1
	}
	if isBarrier(fields, (*guard).positionX, (*guard).positionY) {
		(*guard).positionX, (*guard).positionY = orPositionX, orPositionY
		(*guard).direction = decideDirection((*guard).direction)
	}
}

func printRuneGrid(grid [][]rune) {
	// Iterate over each row
	for _, row := range grid {
		for _, ch := range row {
			// Print each rune as a character
			fmt.Print(string(ch)) // convert rune to string for printing
		}
		fmt.Println() // Move to the next line after each row
	}
}

func main() {
	fields, guard := parseInput("input.txt")
	for !isOutside(fields, guard.positionX, guard.positionY) {
		fields[guard.positionY][guard.positionX] = visitedField
		moveGuard(fields, &guard)
	}
	fmt.Println(countVisited(fields))
	printRuneGrid(fields)

}
