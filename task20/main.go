package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFileTo2DSlice converts a file into a 2D slice of integers and returns the positions of 'S' and 'E'
func ReadFileTo2DSlice(filePath string) ([][]int, [2]int, [2]int, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, [2]int{}, [2]int{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var grid [][]int
	var startPos, endPos [2]int
	startFound, endFound := false, false

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		var rowSlice []int

		for col, char := range line {
			switch char {
			case '#':
				rowSlice = append(rowSlice, -2) // Wall
			case '.':
				rowSlice = append(rowSlice, -1) // Empty space
			case 'S':
				rowSlice = append(rowSlice, 0) // Start position
				startPos = [2]int{row, col}
				startFound = true
			case 'E':
				rowSlice = append(rowSlice, -1) // End position
				endPos = [2]int{row, col}
				endFound = true
			default:
				return nil, [2]int{}, [2]int{}, fmt.Errorf("invalid character '%c' in file", char)
			}
		}

		grid = append(grid, rowSlice)
		row++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, [2]int{}, [2]int{}, fmt.Errorf("failed to read file: %w", err)
	}

	if !startFound || !endFound {
		return nil, [2]int{}, [2]int{}, fmt.Errorf("missing 'S' or 'E' in the input file")
	}

	return grid, startPos, endPos, nil
}

func DecideNewValue(valueStart, valueTarget int) int {
	if valueTarget == -2 {
		return -2
	}
	if valueTarget == -1 {
		return valueStart + 1
	}
	return min(valueStart+1, valueTarget)
}
func BellmanFordIteration(grid [][]int) bool {
	changed := false
	for rowIndex := range len(grid) - 1 {
		for columnIndex := range len(grid[rowIndex]) - 1 {
			value := grid[rowIndex][columnIndex]
			if value >= 0 {
				grid[rowIndex+1][columnIndex] = DecideNewValue(value, grid[rowIndex+1][columnIndex])
				grid[rowIndex-1][columnIndex] = DecideNewValue(value, grid[rowIndex-1][columnIndex])
				grid[rowIndex][columnIndex+1] = DecideNewValue(value, grid[rowIndex][columnIndex+1])
				grid[rowIndex][columnIndex-1] = DecideNewValue(value, grid[rowIndex][columnIndex-1])
			}
		}
	}
	return changed
}
func main() {
	filePath := "maze.txt" // Replace with your file path

	// Call the function to read the file and create the 2D slice
	grid, _, endPos, err := ReadFileTo2DSlice(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	iterations := 100
	for range iterations {
		BellmanFordIteration(grid)
	}
	fmt.Println(grid[endPos[0]][endPos[1]])
}
