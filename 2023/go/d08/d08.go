package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	aocMap, path := getData("assets/input1")

	fmt.Println(aocMap)
	fmt.Println(path)

	steps := followMap(aocMap, path, calcStarts(aocMap))
	fmt.Println(steps)
}

func getData(filename string) (map[string][]string, string) {
	fileData, _ := os.ReadFile(filename)

	fileArr := strings.Split(string(fileData), "\n")
	fileArr = fileArr[:len(fileArr)-1]

	mapPath := fileArr[0]

	aocMap := make(map[string][]string);

	for i := 2; i < len(fileArr) ; i++ {
		lineSplit := strings.Split(fileArr[i], " = ")

		src := lineSplit[0]

		lineSplit[1] = strings.ReplaceAll(lineSplit[1], ")", "")
		lineSplit[1] = strings.ReplaceAll(lineSplit[1], "(", "")
		dst := strings.Split(lineSplit[1], ", ")

		aocMap[src] = dst
	}

	return aocMap, mapPath
}

func calcStarts(aocMap map[string][]string) []string {
	starts := make([]string, 0)

	for key := range aocMap {
		if (endsWith(key, 'A')) {
			starts = append(starts, key)
		}
	}

	return starts
}

func endsWith(src string, final byte) bool {
	last := len(src) - 1

	return src[last] == final
}

func followMap(aocMap map[string][]string, path string, starts []string) int {
	fmt.Println(starts)

	pathMap := make(map[byte]int)
	pathMap['L'] = 0
	pathMap['R'] = 1

	pathIndex := 0
	currentPlaces := starts

	stepsToSolve := make([]int, len(starts))

	steps := 0
	for i := 0; i < len(starts); i++ {
		steps = 0

		for !endsWith(currentPlaces[i], 'Z') {
			currentPlace := currentPlaces[i]
			currentPlaces[i] = aocMap[currentPlace][pathMap[path[pathIndex]]]

			pathIndex++
			if pathIndex == len(path) {
				pathIndex = 0
			}

			steps++
		}

		stepsToSolve[i] = steps
	}

	fmt.Printf("Steps to solve: %v\n", stepsToSolve)

	return lcm(stepsToSolve)
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	return a
}

func lcm(numbers []int) int {
	result := numbers[0] * numbers[1] /gcd(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		result = lcm([]int{result, numbers[i]})
	}

	return result
}
