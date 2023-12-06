package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time, distance int
}

func main() {
	races := GetData("assets/input1")

	fmt.Println(races)

	combinations := getCombination(races)

	fmt.Printf("Combinations: %v\n", combinations)
}

func GetData(filename string) []race {
	fileData, _ := os.ReadFile(filename)

	fileContent := strings.Split(strings.TrimSpace(string(fileData)), "\n")

	re_inside_space := regexp.MustCompile(`[\sp{Zs}]{2,}`)

	times_ := strings.TrimSpace(fileContent[0])
	times_ = strings.Split(times_, ": ")[1]

	// times_ = re_inside_space.ReplaceAllString(times_, " ") // solution 1
	times_ = re_inside_space.ReplaceAllString(times_, "") // solution 2
	times := strings.Split(strings.TrimSpace(times_), " ")

	distances_ := strings.TrimSpace(fileContent[1])
	distances_ = strings.Split(distances_, ": ")[1]
	// distances_ = re_inside_space.ReplaceAllString(distances_, " ") // solution 1
	distances_ = re_inside_space.ReplaceAllString(distances_, "") // solution 2
	distances := strings.Split(strings.TrimSpace(distances_), " ")

	fmt.Println(times)
	fmt.Println(distances)

	len_races := len(times)
	races := make([]race, len_races)

	for i := 0; i < len_races; i++ {
		time, _ := strconv.ParseInt(times[i], 10, 64)
		distance, _ := strconv.ParseInt(distances[i], 10, 64)

		races[i].time = int(time)
		races[i].distance = int(distance)
	}

	return races
}

func getCombination(races []race) int {
	combination := countChances(races[0])

	for i := 1; i < len(races); i++ {
		combination *= countChances(races[i])
	}

	return combination
}

func countChances(r race) int {
	chances := 0

	for speed := 1; speed < r.time; speed++ {
		time := r.time - speed

		finalDistance := speed * time


		if finalDistance > r.distance {
			chances++
		}
	}

	fmt.Printf("Chances: %v\n", chances)
	return chances
}
