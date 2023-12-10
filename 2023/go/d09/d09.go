package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getData("assets/input1")

	fmt.Printf("History: %v\n", reports)
	fmt.Printf("Sum next: %v\n", sumNext(reports))
	fmt.Printf("Sum prev: %v\n", sumPrev(reports))
}

func getData(filename string) [][]int {
	fileData, _ := os.ReadFile(filename)

	fileContent := strings.Split(string(fileData), "\n")
	fileContent = fileContent[:len(fileContent)-1]

	reports := make([][]int, len(fileContent))

	for i, row := range fileContent {
		values := strings.Split(row, " ")

		reports[i] = make([]int, len(values))
		for j, val := range values {
			valInt, _ := strconv.ParseInt(val, 10, 64)
			reports[i][j] = int(valInt)
		}
	}

	return reports
}

func sumPrev(reports [][]int) int {
	sum := 0

	for _, report := range reports {
		sum += calcPrev(report)
	}

	return sum
}

func calcPrev(report []int) int {
	stepMatrix := createStepMatrix(report)
	stepMatrixLayers := len(stepMatrix)

	fmt.Println(stepMatrix)

	stepMatrix[stepMatrixLayers-1] = append(stepMatrix[stepMatrixLayers-1], 0)
	for i := stepMatrixLayers-2; i >= 0; i-- {
		step := stepMatrix[i+1][0]

		firstValue := stepMatrix[i][0]

		stepMatrix[i] = append([]int{firstValue-step}, stepMatrix[i]...)
	}

	return stepMatrix[0][0]
}

func sumNext(reports [][]int) int {
	sum := 0

	for _, report := range reports {
		fmt.Println(calcNext(report))
		sum += calcNext(report)
	}

	return sum
}

func calcNext(report []int) int {
	stepMatrix := createStepMatrix(report)
	stepMatrixLayers := len(stepMatrix)

	fmt.Println(stepMatrix)

	stepMatrix[stepMatrixLayers-1] = append(stepMatrix[stepMatrixLayers-1], 0)
	for i := stepMatrixLayers-2; i >= 0; i-- {
		currentLen := len(stepMatrix[i])
		stepLen := len(stepMatrix[i+1])
		step := stepMatrix[i+1][stepLen-1]

		lastValue := stepMatrix[i][currentLen-1]

		stepMatrix[i] = append(stepMatrix[i], lastValue+step)
	}

	firstLayerLen := len(stepMatrix[0])
	return stepMatrix[0][firstLayerLen-1]
}

func createStepMatrix(report []int) [][]int {
	stepMatrix := make([][]int, 1)

	stepMatrix[0] = report

	isAllZero := false
	stepDeep := 0
	for !isAllZero {
		isAllZero = true
		currentStep := make([]int, 0)

		currentStepArray := stepMatrix[stepDeep]
		for i := 0; i < len(stepMatrix[stepDeep]) - 1; i++ {
			currentStep = append(currentStep, currentStepArray[i+1] - currentStepArray[i])
		}
		stepMatrix = append(stepMatrix, currentStep)
		stepDeep++

		for _, value := range stepMatrix[stepDeep] {
			if value != 0 {
				isAllZero = false
			}
		}
	}

	stepMatrix = append(stepMatrix, make([]int, len(stepMatrix[stepDeep])-1))

	return stepMatrix
}
