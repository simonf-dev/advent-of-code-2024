
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

func parseStones(filePath string) []int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	stones := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		for _, value := range strings.Split(text, " ") {
			stoneNumber, err := strconv.Atoi(value)
			check(err)
			stones = append(stones, stoneNumber)
		}
	}
	return stones
}

func convertStone(stoneNumber int) []int {
	newStones := make([]int, 0)
	numberStr := strconv.Itoa(stoneNumber)
	if numberStr == "0" {
		newStones = append(newStones, 1)
	} else if len(numberStr)%2 == 0 {
		mid := len(numberStr) / 2
		firstHalfStr := numberStr[:mid]
		secondHalfStr := numberStr[mid:]

		firstHalf, err := strconv.Atoi(firstHalfStr)
		check(err)
		secondHalf, err := strconv.Atoi(secondHalfStr)
		check(err)

		newStones = append(newStones, firstHalf, secondHalf)
	} else {
		newStones = append(newStones, 2024*stoneNumber)
	}
	return newStones
}

func updateStones(blinks int, stones []int) int {
	for index := range blinks {
		newStones := make([]int, 0)
		for _, stoneNumber := range stones {
			newStones = append(newStones, convertStone(stoneNumber)...)
		}
		stones = newStones
		fmt.Println("index", index)
	}
	return len(stones)
}
func main() {
	stones := parseStones("input.txt")
	fmt.Println(updateStones(75, stones))
}
