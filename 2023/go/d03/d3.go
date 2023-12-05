package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type aoc_number struct {
	value, x, y, size int;
	found bool;
}

func main() {
	input := getData("assets/input2")

	fmt.Println(input)
	fmt.Println("len :", len(input))

	res := findNumbers(input)

	fmt.Println("res: ", res)
}

func getData(file_path string) []string {
	file_data, _ := os.ReadFile(file_path)

	res := strings.Split(string(file_data), "\n")
	res_len := len(res)

	return res[:res_len-1]
}

func findNumbers(input []string) int {
	rows := len(input)
	cols := len(input[0])

	// numbers := make([]int, rows * cols)
	sum := 0

	numList := createNumbersList(input)
	fmt.Println(numList)

	for i := 0; i < rows ; i++ {
		for j := 0; j < cols ; j++ {
			if isSymbol(input[i][j]) {
				sum += getNeightborsSum(&numList, j, i, rows, cols)
			}
		}
	}
	fmt.Println(numList)

	return sum
}

func isSymbol(val byte) bool {
	num := int(val)
	return (num < int('0') || num > int('9')) && num != int('.')
}

func createNumbersList(input []string) []aoc_number {
	rows := len(input)
	cols := len(input[0])

	numList := make([]aoc_number, 0)

	for i := 0 ; i < rows ; i++ {
		var start int = -1
		numStack := ""

		for j := 0 ; j < cols ; j++ {
			char := input[i][j]
			if isNumber(char) {
				numStack += string([]byte {char})

				if start == -1 {
					start = j
				}
			} else if start != -1 {
				num, _ := strconv.ParseInt(numStack, 10, 64)

				if num > 0 {
					new_number := aoc_number {
						value: int(num),
						x: start,
						y: i,
						size: len(numStack),
					}

					numList = append(numList, new_number)
				}

				numStack = ""
				start = -1
			}
		}
		if start != -1 {
			num, _ := strconv.ParseInt(numStack, 10, 64)

			if num > 0 {
				new_number := aoc_number {
					value: int(num),
					x: start,
					y: i,
					size: len(numStack),
				}

				numList = append(numList, new_number)
			}

			numStack = ""
			start = -1
		}
	}

	return numList
}

func isNumber(val byte) bool {
	return val >= '0' && val <= '9'
}

func getNeightborsSum(numbers *[]aoc_number, x int, y int, rows int, cols int) int {

	coords := [][2]int {
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	};

	numsToSum := make([]int, 0)
	for i := 0; i < len(coords); i++ {
		offset_y := coords[i][0] + y
		offset_x := coords[i][1] + x

		is_valid := offset_y >= 0 && offset_y < cols &&
			offset_x >= 0 && offset_x < rows

		if is_valid {
			for k := 0; k < len(*numbers) ; k++ {
				val := &(*numbers)[k]

				for c := val.x ; c < val.x + val.size ; c++ {
					isNeightbor := c == offset_x && val.y == offset_y && !val.found
					if isNeightbor {
						numsToSum = append(numsToSum, val.value)
						val.found = true
						break
					}
				}

			}
		}
	}

	sum := 0
	// solution 2
	if len(numsToSum) == 2 {
		sum = numsToSum[0] * numsToSum[1]
	}

	return sum
}
