package main

import (
	"fmt"
	"strings"
	// "bufio"
	// "io"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args[1])
	input := readFile(args[1], len(args) > 2)

	var sum = 0
	for i := 0; i < len(input); i++ {
		var result = findNumber(input[i])
		sum += result
	}

	fmt.Printf("Total sum: %d\n", sum)
}

func readFile(file_name string, should_preprocess bool) []string {
	data, err := os.ReadFile(file_name)

	if (err != nil) {
		fmt.Println("Erro!")
		return nil
	}

	input := string(data)

	if should_preprocess {
		input = preProcessString(input)
	}

	input_arr := strings.Split(input, "\n")

	return input_arr
}

// Probably the wors solution possible.
func preProcessString(source string) string {
	source = strings.ReplaceAll(source, "one", "o1e")
	source = strings.ReplaceAll(source, "two", "t2o")
	source = strings.ReplaceAll(source, "three", "t3e")
	source = strings.ReplaceAll(source, "four", "f4r")
	source = strings.ReplaceAll(source, "five", "f5e")
	source = strings.ReplaceAll(source, "six", "s6x")
	source = strings.ReplaceAll(source, "seven", "s7n")
	source = strings.ReplaceAll(source, "eight", "e8t")
	source = strings.ReplaceAll(source, "nine", "n9e")

	return source
}

func findNumber(line string) int {
	var first_number_done, last_number_done = false, false
	var first, last int = 0, 0

	line_size := len(line)
	for i := 0; i < line_size ; i++ {
		var first_char = line[i]
		var last_char = line[line_size-i-1]

		if !first_number_done && isDigit(first_char) {
			first = int(first_char - 0x30)
			first_number_done = true
		}

		if !last_number_done && isDigit(last_char) {
			last = int(last_char - 0x30)
			last_number_done = true
		}

		if first_number_done && last_number_done {
			break
		}
	}


	var final_number = first * 10 + last

	return final_number
}

func isDigit(char byte) bool {
	return char >= 0x30 && char <= 0x39
}
