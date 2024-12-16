package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
	"math"
	"errors"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parse_file(filePath string) [2][]int {
	dat, err := os.ReadFile(filePath)
	check(err)
	dat_str := string(dat)
	lines := strings.Split(dat_str, "\n")
	var arr [2][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		for i := range 2 {
			converted_value, err := strconv.Atoi(fields[i])
			check(err)
			arr[i] = append(arr[i], converted_value)
		}
	} 
	return arr
}

func sort_list (entry_list []int) {
	for i := range len(entry_list) {
		for j := range len(entry_list) - i - 1 {
			if entry_list[j] > entry_list[j+1]{
				tmp_var := entry_list[j+1]
				entry_list[j+1] = entry_list[j]
				entry_list[j] = tmp_var
			}
		}
	}
}

func count_distance(value_array [2][]int) (int, error) {
	var err error = nil
	length_fst := len(value_array[0])
	length_snd := len(value_array[1])
	var distance int = 0
	if length_fst != length_snd {
		err = errors.New("length of first and second array isn't the same")
	}
	for i := range length_fst {
		distance += int(math.Abs(float64(value_array[0][i] - value_array[1][i])))
	} 	
	return distance, err
}

func count_similarity(value_array [2][]int) int {
	map_snd := make(map[int]int)
	for _, item := range value_array[1] {
		map_snd[item] = map_snd[item] + 1
	}
	result := 0
	for _, value_fst := range value_array[0] {
		value_snd := map_snd[value_fst]
		result = result + value_snd * value_fst
	}
	return result
	
}

func main() {
	value_array := parse_file("entry_list.csv")
	sort_list(value_array[0])
	sort_list(value_array[1])
	val, err := count_distance(value_array)
	check(err)
	fmt.Println(val)
	fmt.Println(count_similarity(value_array))
}
