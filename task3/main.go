
package main
import (
	"os"
	"regexp"
	"fmt"
	"strconv"
	"errors"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func parse_input(filePath string) string {
	dat, err := os.ReadFile(filePath)
	check(err)
	dat_str := string(dat)
	return dat_str
}



func main() {
	input_str := parse_input("input.txt")
	r, _ := regexp.Compile("mul\\(\\d*,\\d*\\)|don't\\(\\)|do\\(\\)")
	mul_list := r.FindAllString(input_str, -1)
	r, _ = regexp.Compile("\\d+")
	result := 0
	enabled := true
	for _, value := range mul_list {
		if value == "do()" {
			enabled = true
		} else if value == "don't()" {
			enabled = false
		} else if enabled {
			decimals := r.FindAllString(value, -1)
			if len(decimals) != 2 {
			    fmt.Println(decimals)
				panic(errors.New("incorrect mul structure"))
			}
			first_dec, err := strconv.Atoi(decimals[0])
			check(err)
			snd_dec, err := strconv.Atoi(decimals[1])
			check(err)
	        result += snd_dec * first_dec
        }
	}
	fmt.Println(result)
}
