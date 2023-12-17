package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type maze struct {
	mazeMap [][]byte;
	start [2]int;
}

var pipePaths = map[byte][2][2]int {
	'-': {
		{ -1, 0 },
		{  1, 0 },
	},
	'|': {
		{ 0, -1 },
		{ 0,  1 },
	},
	'7': {
		{ -1, 0 },
		{  0, 1 },
	},
	'J': {
		{  0, -1 },
		{ -1,  0 },
	},
	'L': {
		{ 1,  0 },
		{ 0, -1 },
	},
	'F': {
		{ 1, 0 },
		{ 0, 1 },
	},
	'.': {
		{ 0, 0 },
		{ 0, 0 },
	},
}

func main() {
	aocMap := getData("assets/input1")

	fmt.Println(aocMap)

	fmt.Println("Hello, World!")

	walk(aocMap)
}

func getData(filename string) maze {
	fileData, _ := os.ReadFile(filename)

	fileContent := strings.Split(string(fileData), "\n")
	fileContent = fileContent[:len(fileContent)-1]

	var aocMap maze

	aocMap.mazeMap = make([][]byte, len(fileContent))
	for i, row := range fileContent {
		aocMap.mazeMap[i] = make([]byte, len(fileContent[0]))
		for j, val := range row {
			if val == 'S' {
				aocMap.start = [2]int{j, i}
				aocMap.mazeMap[i][j] = 'J'
				continue
			}

			aocMap.mazeMap[i][j] = byte(val)
		}
	}

	return aocMap
}

func walk(aocMap maze) {
	from := aocMap.start
	currentX := aocMap.start[0]
	currentY := aocMap.start[1]

	currentPipe := aocMap.mazeMap[currentY][currentX]
	// fmt.Printf("(%v, %v): %c\n", currentX, currentY, currentPipe)

	newX := currentX + pipePaths[currentPipe][0][0]
	newY := currentY + pipePaths[currentPipe][0][1]

	if newX == from[0] && newY == from[1] {
		newX = currentX + pipePaths[currentPipe][1][0]
		newY = currentY + pipePaths[currentPipe][1][1]
	}

	from = [2]int{currentX, currentY}
	currentX = newX
	currentY = newY

	farthest := 0

	distMap := make([][]int, len(aocMap.mazeMap))

	for i := 0; i < len(aocMap.mazeMap); i++ {
		distMap[i] = make([]int, len(aocMap.mazeMap[0]))
	}

	distance := 1
	for true {
		currentPipe = aocMap.mazeMap[currentY][currentX]

		newX = currentX + pipePaths[currentPipe][0][0]
		newY = currentY + pipePaths[currentPipe][0][1]

		currentFarthest := int(math.Abs(float64(aocMap.start[0] - currentX)))
		currentFarthest += int(math.Abs(float64(aocMap.start[1] - currentY)))

		// fmt.Printf("[%v, %v] - [%v, %v]\n", aocMap.start[0], aocMap.start[1], currentX, currentY);
		// fmt.Printf("(%v, %v): %c => %v\n", currentX, currentY, currentPipe, currentFarthest)
		farthest = max(farthest, currentFarthest)

		if distMap[currentY][currentX] != 0 {
			distMap[currentY][currentX] = min(distance, distMap[currentY][currentX])
		} else {
			distMap[currentY][currentX] = distance
		}

		if newX == from[0] && newY == from[1] {
			newX = currentX + pipePaths[currentPipe][1][0]
			newY = currentY + pipePaths[currentPipe][1][1]
		}

		from = [2]int{currentX, currentY}
		currentX = newX
		currentY = newY

		if from == aocMap.start {
			break
		}
		distance++
	}

	printDistMap(distMap)

	from = aocMap.start
	currentX = aocMap.start[0]
	currentY = aocMap.start[1]

	currentPipe = aocMap.mazeMap[currentY][currentX]
	// fmt.Printf("(%v, %v) %c\n", currentX, currentY, currentPipe)

	newX = currentX + pipePaths[currentPipe][1][0]
	newY = currentY + pipePaths[currentPipe][1][1]

	if newX == from[0] && newY == from[1] {
		newX = currentX + pipePaths[currentPipe][0][0]
		newY = currentY + pipePaths[currentPipe][0][1]
	}

	from = [2]int{currentX, currentY}
	currentX = newX
	currentY = newY

	distance = 1
	for true {
		currentPipe = aocMap.mazeMap[currentY][currentX]

		newX = currentX + pipePaths[currentPipe][0][0]
		newY = currentY + pipePaths[currentPipe][0][1]

		currentFarthest := int(math.Abs(float64(aocMap.start[0] - currentX)))
		currentFarthest += int(math.Abs(float64(aocMap.start[1] - currentY)))

		// fmt.Printf("[%v, %v] - [%v, %v]\n", aocMap.start[0], aocMap.start[1], currentX, currentY);
		// fmt.Printf("(%v, %v): %c => %v\n", currentX, currentY, currentPipe, currentFarthest)
		farthest = max(farthest, currentFarthest)

		if distMap[currentY][currentX] != 0 {
			distMap[currentY][currentX] = min(distance, distMap[currentY][currentX])
		} else {
			distMap[currentY][currentX] = distance
		}

		if newX == from[0] && newY == from[1] {
			newX = currentX + pipePaths[currentPipe][1][0]
			newY = currentY + pipePaths[currentPipe][1][1]
		}

		from = [2]int{currentX, currentY}
		currentX = newX
		currentY = newY

		if from == aocMap.start {
			break
		}
		distance++
	}

	fmt.Println(farthest)

	distMap[aocMap.start[1]][aocMap.start[0]] = 0

	printDistMap(distMap)

	fmt.Println(findMax(distMap))
}

func printDistMap(distMap [][]int) {
	for _, row := range distMap {
		for _, char := range row {
			fmt.Printf("%04v ", char)
		}
		fmt.Println()
	}
}

func findMax(distMap [][]int) int {
	maxValue := distMap[0][0]

	for _, row := range distMap {
		for _, val := range row {
			maxValue = max(maxValue, val)
		}
	}

	return maxValue
}
