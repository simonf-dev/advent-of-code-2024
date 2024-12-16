package main
import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func parse_input(filePath string) [][]int {
	dat, err := os.ReadFile(filePath)
	check(err)
	dat_str := string(dat)
	lines := strings.Split(dat_str, "\n")
	var arr [][]int
	for _, line := range lines {
	        fields := strings.Fields(line)
	        if len(fields) < 2 {
                continue
	        }
	        var line []int
	        for i := range len(fields) {
                converted_value, err := strconv.Atoi(fields[i])
                check(err)
                line = append(line, converted_value)
	        }
	        arr = append(arr, line)
	} 
	return arr
}


func checkLine(parsedLine []int) bool {
	if len(parsedLine) < 2 {
		return true
	}
	increasing := parsedLine[0] < parsedLine[1]
	for i := 0; i < len(parsedLine) - 1; i++ {
		absDifference := int(math.Abs(float64(parsedLine[i] - parsedLine[i+1])))
		if (parsedLine[i] < parsedLine[i+1]) != increasing || (absDifference < 1 || absDifference > 3) {
			return false
		}
	}
	return true
}
func main() {
	parsedList := parse_input("input.csv")
	correctLines := 0
	for _, line := range parsedList {
		if checkLine(line) {
			correctLines += 1
			continue
		}
		for index := range len(line) {
			startPart := make([]int, index)
			endPart := make([]int, len(line)-index-1)
			copy(endPart, line[index+1:])
			copy(startPart, line[:index])
			if checkLine(append(startPart, endPart...)) {
				correctLines += 1
				break
			}
		}
	}
	fmt.Println(correctLines)
}
