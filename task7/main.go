package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func concatNumbers(num1, num2 int) int {

	strResult := strconv.Itoa(num1) + strconv.Itoa(num2)
	result, err := strconv.Atoi(strResult)
	check(err)
	return result
}
func parseRules(filePath string) map[int][][]int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	operationInfo := make(map[int][][]int)
	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), ":")
		if len(lineSlice) != 2 {
			panic("There should be exactly one sum value")
		}
		resultValue, err := strconv.Atoi(lineSlice[0])
		check(err)
		lineSlice[1] = strings.TrimSpace(lineSlice[1])
		splittedLine := strings.Split(lineSlice[1], " ")
		fmt.Println(splittedLine)
		operationValues := make([]int, len(splittedLine))
		for index, strValue := range splittedLine {
			resultValue, err := strconv.Atoi(strValue)
			check(err)
			operationValues[index] = resultValue
		}
		_, contains := operationInfo[resultValue]
		if contains {
			operationInfo[resultValue] = append(operationInfo[resultValue], operationValues)
		} else {
			operationInfo[resultValue] = make([][]int, 1)
			operationInfo[resultValue][0] = operationValues
		}
	}
	return operationInfo
}

func checkCombinations(actualValue, resultValue, currentIndex int, values []int) bool {
	if actualValue == resultValue && currentIndex == len(values) {
		return true
	}
	if currentIndex >= len(values) {
		return false
	}
	return checkCombinations(actualValue*values[currentIndex], resultValue, currentIndex+1, values) || checkCombinations(actualValue+values[currentIndex], resultValue, currentIndex+1, values) || checkCombinations(concatNumbers(actualValue, values[currentIndex]), resultValue, currentIndex+1, values)
}
func main() {
	output := parseRules("input.txt")
	fmt.Println(len(output))
	returnOutput := 0
	for resultValue, valueSlice := range output {
		for _, values := range valueSlice {
			if checkCombinations(values[0], resultValue, 1, values) {
				returnOutput += resultValue
				// fmt.Println(resultValue)
			}
		}
	}
	fmt.Println(returnOutput)
}
