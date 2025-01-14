package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const FIELD_WIDE = 101
const FIELD_TALL = 103

type Position struct {
	X, Y int
}
type Robot struct {
	position, velocity Position
}

type Quadrant struct {
	A, B, C, D int
}

func (robot *Robot) MoveRobot() {
	robot.position.X = (robot.position.X + robot.velocity.X) % FIELD_WIDE
	if robot.position.X < 0 {
		robot.position.X = FIELD_WIDE + robot.position.X
	}
	robot.position.Y = (robot.position.Y + robot.velocity.Y) % FIELD_TALL
	if robot.position.Y < 0 {
		robot.position.Y = FIELD_TALL + robot.position.Y
	}
}

func (robot Robot) DecideQuadrant(quadrant *Quadrant) {
	halfWide := FIELD_WIDE / 2
	halfTall := FIELD_TALL / 2
	if robot.position.X < halfWide && robot.position.Y < halfTall {
		quadrant.A += 1
	} else if robot.position.X > halfWide && robot.position.Y < halfTall {
		quadrant.B += 1
	} else if robot.position.X < halfWide && robot.position.Y > halfTall {
		quadrant.C += 1
	} else if robot.position.X > halfWide && robot.position.Y > halfTall {
		quadrant.D += 1
	}
}

func (quadrant Quadrant) IsSymmetric() bool {
	return quadrant.A == quadrant.B && quadrant.C == quadrant.D
}
func (quadrant Quadrant) CountIndex() int {
	return quadrant.A * quadrant.B * quadrant.C * quadrant.D
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string, re *regexp.Regexp) Robot {
	matches := re.FindStringSubmatch(line)
	if len(matches) != 5 {
		panic(fmt.Sprintf("Unexpected format: %s", line))
	}
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	velX, _ := strconv.Atoi(matches[3])
	velY, _ := strconv.Atoi(matches[4])
	return Robot{position: Position{X: x, Y: y}, velocity: Position{X: velX, Y: velY}}
}

func parseRobots(filePath string) []Robot {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	robots := make([]Robot, 0)
	for scanner.Scan() {
		text := scanner.Text()
		valueRegex := regexp.MustCompile(`p\=(\d+),(\d+).*v=(-?\d+),(-?\d+)`)
		robots = append(robots, parseLine(text, valueRegex))
	}
	return robots
}

func ExecuteMoves(robots []Robot, seconds int) {
	for second := range seconds {
		for robotIndex := range robots {
			robots[robotIndex].MoveRobot()
		}
		if second > 7566 {
			writeRobotPositionsAsMatrixToFile(robots, FIELD_WIDE, FIELD_TALL, "output.txt", second)
		}

	}
}

func CountQuadrants(robots []Robot) Quadrant {
	quadrant := Quadrant{}
	for robotIndex := range robots {
		robots[robotIndex].DecideQuadrant(&quadrant)
	}
	return quadrant
}

func writeRobotPositionsAsMatrixToFile(robots []Robot, width, height int, filename string, iteration int) error {
	// Create a matrix to represent the grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.' // Initialize the grid with '.' for empty spaces
		}
	}

	// Place robots on the grid
	for _, robot := range robots {
		// Ensure the robot's position is within bounds
		if robot.position.X >= 0 && robot.position.X < width &&
			robot.position.Y >= 0 && robot.position.Y < height {
			grid[robot.position.Y][robot.position.X] = rune('X') // Label robots as 'X'
		}
	}

	// Open the file in append mode or create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Write the iteration header to the file
	header := fmt.Sprintf("Iteration %d\n", iteration)
	if _, err := file.WriteString(header); err != nil {
		return fmt.Errorf("failed to write header to file: %w", err)
	}

	// Write the grid to the file
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if _, err := file.WriteString(fmt.Sprintf("%c", grid[i][j])); err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
		}
		if _, err := file.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to write newline to file: %w", err)
		}
	}

	// Write a separator line
	if _, err := file.WriteString("\n---\n\n"); err != nil {
		return fmt.Errorf("failed to write separator to file: %w", err)
	}

	return nil
}
func main() {
	robots := parseRobots("input.txt")
	ExecuteMoves(robots, 7600)
	fmt.Println(CountQuadrants(robots))

}
