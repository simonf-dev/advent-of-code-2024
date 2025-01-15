package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	adv Operation = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetComboValue(register map[rune]int, value int) int {
	switch value {
	case 0, 1, 2, 3:
		return value
	case 4:
		return register['A']
	case 5:
		return register['B']
	case 6:
		return register['C']
	default:
		panic("invalid combo operator")
	}
}

func PerformAdv(value int, register map[rune]int) {
	fmt.Println("Operation: A=", register['A'], "/ 2^", value)
	register['A'] = register['A'] / (1 << value)
	fmt.Println("Register A: ", register['A'])
}

func PerformBdv(value int, register map[rune]int) {
	fmt.Println("Operation: B=", register['A'], "/ 2^", value)
	register['B'] = register['A'] / (1 << value)
	fmt.Println("Register B: ", register['B'])
}

func PerformCdv(value int, register map[rune]int) {
	fmt.Println("Operation: C=", register['A'], "/ 2^", value)
	register['C'] = register['A'] / (1 << value)
	fmt.Println("Register C: ", register['C'])
}

func PefrormBxl(value int, register map[rune]int) {
	fmt.Println("Operation: B=", register['B'], "^", value)
	register['B'] = register['B'] ^ value
	fmt.Println("Register B: ", register['B'])
}

func PerformBst(value int, register map[rune]int) {
	fmt.Println("Operation: B=", value, "%", 8)
	register['B'] = value % 8
	fmt.Println("Register B: ", register['B'])
}

func PerformBxc(register map[rune]int) {
	fmt.Println("Operation: B=", register['B'], "^", register['C'])
	register['B'] = register['B'] ^ register['C']
	fmt.Println("Register B: ", register['B'])
}

func PerformExecution(pointer int, program []int, register map[rune]int) string {
	output := make([]int, 0)
	for pointer < len(program) {
		if pointer%2 != 0 {
			panic("Pointer has to be even")
		}
		operation := Operation(program[pointer])
		value := program[pointer+1]
		pointer += 2
		switch operation {
		case adv:
			PerformAdv(GetComboValue(register, value), register)
		case bxl:
			PefrormBxl(value, register)
		case bst:
			PerformBst(GetComboValue(register, value), register)
		case jnz:
			if register['A'] != 0 {
				fmt.Println("Jumping to", value)
				pointer = value
			}
		case bxc:
			PerformBxc(register)
		case out:
			fmt.Println("Appending", GetComboValue(register, value)%8)
			output = append(output, GetComboValue(register, value)%8)
		case bdv:
			PerformBdv(GetComboValue(register, value), register)
		case cdv:
			PerformCdv(GetComboValue(register, value), register)
		}
	}
	result := strings.Join(func(nums []int) []string {
		strs := make([]string, len(nums))
		for i, n := range nums {
			strs[i] = strconv.Itoa(n)
		}
		return strs
	}(output), "")
	return result
}

func ParseCommaSeparatedFile(filepath string) ([]int, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Use a scanner to read the file (useful for large files)
	scanner := bufio.NewScanner(file)

	var nums []int
	for scanner.Scan() {
		// Get the current line
		line := strings.TrimSpace(scanner.Text())

		// Split the line by commas
		parts := strings.Split(line, ",")

		// Convert each part to an integer
		for _, part := range parts {
			part = strings.TrimSpace(part) // Trim any stray spaces around the number
			if part == "" {                // Skip empty values (e.g., trailing commas)
				continue
			}

			// Convert the string to an integer
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number %q: %w", part, err)
			}
			nums = append(nums, num)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return nums, nil
}
func main() {
	program, err := ParseCommaSeparatedFile("input")
	check(err)
	registers := make(map[rune]int)
	registers['A'] = 7 * 20

	registers['B'] = 0
	registers['C'] = 0
	result := PerformExecution(0, program, registers)
	fmt.Println(result)

}
