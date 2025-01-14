package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Equation struct {
	A, B, Constant int
}

type EquationCombination struct {
	X, Y Equation
}

func (eq *Equation) SolveForA(a int) int {
	return (eq.Constant - a*eq.A) / eq.B
}

func (eq *Equation) SolveForB(b int) int {
	if (eq.Constant-b*eq.B)%eq.A != 0 {
		return -1
	}
	return (eq.Constant - b*eq.B) / eq.A
}

func (eq *Equation) Solve(a, b int) bool {
	return a*eq.A+b*eq.B == eq.Constant
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solveEquations(combination EquationCombination) int {
	parsedEquation := Equation{B: combination.Y.B*combination.X.A - combination.X.B*combination.Y.A,
		Constant: combination.Y.Constant*combination.X.A - combination.X.Constant*combination.Y.A,
	}
	if parsedEquation.B == 0 || parsedEquation.Constant%parsedEquation.B != 0 {
		return 0
	}
	b := parsedEquation.Constant / parsedEquation.B
	a := combination.X.SolveForB(b)
	if a < 0 {
		return 0
	}
	return b + a*3
}

func parseLine(line string, re *regexp.Regexp) (int, int) {
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		panic(fmt.Sprintf("Unexpected format: %s", line))
	}
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}

func parseEquations(filePath string) []EquationCombination {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	equations := make([]EquationCombination, 0)
	for scanner.Scan() {
		text := scanner.Text()
		buttonRe := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
		prizeRe := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
		aX, aY := parseLine(text, buttonRe)
		scanner.Scan()
		text = scanner.Text()
		bX, bY := parseLine(text, buttonRe)
		scanner.Scan()
		text = scanner.Text()
		resultX, resultY := parseLine(text, prizeRe)
		xEquation := Equation{A: aX, B: bX, Constant: resultX + 10000000000000}
		yEquation := Equation{A: aY, B: bY, Constant: resultY + 10000000000000}
		scanner.Scan()
		equations = append(equations, EquationCombination{X: xEquation, Y: yEquation})

	}
	return equations
}

func main() {
	equations := parseEquations("input.txt")
	var tokens int64 = 0
	for _, equation := range equations {
		newTokens := int64(solveEquations(equation))
		fmt.Println("Tokens for equation", equation, newTokens)
		tokens += newTokens
	}
	fmt.Println(tokens)
}
