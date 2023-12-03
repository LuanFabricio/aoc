package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	games, bag := readData("assets/input1")

	fmt.Println(games)
	fmt.Println(bag)

	possible_games := getPossibleGames(games, bag)
	var sum uint = 0

	for i := 0 ; i < len(possible_games) ; i++ {
		sum += possible_games[i]
	}

	fmt.Printf("Sum: %v\n", sum)
	fmt.Printf("Min sum: %v\n", minSetPower(games))
}

type cube struct {
	blue, green, red uint64;
}

func readData(filePath string) ([]cube, cube) {
	var filedata, err = os.ReadFile(filePath)

	if (err != nil) {
		fmt.Println("Erro!")
	}

	var file_arr = strings.Split(string(filedata), "\n")

	games_size := len(file_arr) - 1
	games_max := make([]cube, games_size)
	for i := 0 ; i < len(file_arr) - 1 ; i++ {
		game_str := file_arr[i]
		fmt.Println(game_str)

		game_str_split := strings.Split(game_str, ": ")
		cube_sets := strings.Split(game_str_split[1], "; ")

		cube := cube {}
		game_index_str := (strings.Split(game_str_split[0], " ")[1])
		game_index, _ := strconv.ParseUint(game_index_str, 10, 64)
		for j := 0 ; j < len(cube_sets) ; j++ {
			current_set := getSetCubes(cube_sets[j])
			cube.red = max(cube.red, current_set.red)
			cube.green = max(cube.green, current_set.green)
			cube.blue = max(cube.blue, current_set.blue)
		}

		fmt.Printf("Game index: %v\n", game_index)
		games_max[game_index - 1] = cube
	}

	bag := cube { red: 12, green: 13, blue: 14 }
	return games_max, bag
}

func getSetCubes(game_set string) cube {
	cube_set := cube { blue: 0, red: 0, green: 0}

	cubes_str := strings.Split(game_set, ", ")

	for i := 0 ; i < len(cubes_str) ; i++ {
		fmt.Println(cubes_str[i])

		item_split := strings.Split(cubes_str[i], " ")
		val, _ := strconv.ParseUint(item_split[0], 10, 64)
		var key = item_split[1]

		switch key {
			case "red":
				cube_set.red = val
				break
			case "green":
				cube_set.green = val
				break
			case "blue":
				cube_set.blue = val
				break
			default:
				break
		}
	}

	return cube_set
}

 func getPossibleGames(games []cube, bag cube) []uint {
	games_size := uint(len(games))
	games_id := make([]uint, games_size)
	last_index := 0

	for i := uint(0) ; i < games_size ; i++ {
		if games[i].red <= bag.red && games[i].green <= bag.green && games[i].blue <= bag.blue {
			games_id[last_index] = i+1
			last_index++
		}
	}

	return games_id
}

func minSetPower(games []cube) uint {
	var sum uint = 0

	for i := 0 ; i < len(games) ; i++ {
		sum += uint(games[i].red) * uint(games[i].green) * uint(games[i].blue)
	}

	return sum
}
