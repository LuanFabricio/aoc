package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type almanac struct {
	seeds [][2]int;
	seed_to_soil, soil_to_fertilizer,
	fertilizer_to_water, water_to_light,
	light_to_temperature, temperature_to_humidity,
	humidity_to_location []almanac_map;
}

type almanac_map struct {
	source, dest, arange int
}

var wg sync.WaitGroup

func main() {
	data := getData("assets/input1")

	fmt.Printf("File data: %v\n", data)

	shortest_location := findShortestLocation(data)

	fmt.Println("shortest location: ", shortest_location)
}

func getData(filename string) almanac {
	file_data, _ := os.ReadFile(filename)

	file_content := strings.Split(strings.TrimSpace(string(file_data)), "\n")
	file_content = file_content[:len(file_content)-1]

	// seeds_arr := getSeeds(file_content[0])
	seeds_arr := getSeeds2(file_content[0])

	my_almanac := almanac {
		seeds: seeds_arr,
	}

	for i := 2; i < len(file_content); i++ {
		current_line := file_content[i]

		i++
		line := file_content[i]
		amap_arr := make([]almanac_map, 0)

		for len(line) > 1 {
			amap_arr = append(amap_arr, getMap(line))

			i++
			if i < len(file_content) {
				line = file_content[i]
			} else {
				break
			}
		}

		switch current_line {
		case "seed-to-soil map:":
			my_almanac.seed_to_soil = amap_arr
			break
		case "fertilizer-to-water map:":
			my_almanac.fertilizer_to_water = amap_arr
			break
		case "soil-to-fertilizer map:":
			my_almanac.soil_to_fertilizer = amap_arr
			break
		case "water-to-light map:":
			my_almanac.water_to_light = amap_arr
			break
		case "light-to-temperature map:":
			my_almanac.light_to_temperature = amap_arr
			break
		case "temperature-to-humidity map:":
			my_almanac.temperature_to_humidity = amap_arr
			break
		case "humidity-to-location map:":
			my_almanac.humidity_to_location = amap_arr
			break
		default:
			break
		}
	}

	return my_almanac
}

func getSeeds(line string) []int {
	seeds_str := strings.Split(line, ": ")[1]

	seeds_num := strings.Split(seeds_str, " ")

	fmt.Println(seeds_num)
	seeds_arr := make([]int, len(seeds_num))

	for i, num := range seeds_num {
		val, _ := strconv.ParseInt(num, 10, 64)
		seeds_arr[i] = int(val)
	}

	return seeds_arr
}

func getSeeds2(line string) [][2]int {
	seeds_str := strings.Split(line, ": ")[1]

	seeds_num := strings.Split(seeds_str, " ")

	fmt.Println(seeds_num)
	seeds_arr := make([][2]int, 0)

	for i := 0 ; i < len(seeds_num) ; i++ {
		inits := seeds_num[i]
		i++
		sranges := seeds_num[i]

		init, _ := strconv.ParseInt(inits, 10, 64)
		srange, _ := strconv.ParseInt(sranges, 10, 64)

		seeds_arr = append(seeds_arr, [2]int{ int(init), int(srange) })
	}

	return seeds_arr
}

func getMap(line string) almanac_map {
	amap := strings.Split(line, " ")
	dest_str, source_str, arange_str := amap[0], amap[1], amap[2]

	dest, _ := strconv.ParseInt(dest_str, 10, 64)
	source, _ := strconv.ParseInt(source_str, 10, 64)
	arange, _ := strconv.ParseInt(arange_str, 10, 64)

	return almanac_map {
		dest: int(dest),
		source: int(source),
		arange: int(arange),
	}
}

func findShortestLocation(my_almanac almanac) int {
	max_uint := ^uint(0)
	max_int := int(max_uint >> 1)
	shortest := max_int

	seed_len := len(my_almanac.seeds)
	seed_chan := make([]chan int, seed_len)

	for i := 0; i < seed_len; i++ {

		seed_chan[i] = make(chan int, 1)
		go func (i int) {
			res := findShortestLocationByRange(my_almanac, my_almanac.seeds[i])

			seed_chan[i] <- res
			close(seed_chan[i])
		}(i)

		// Sequential solution
		// init := my_almanac.seeds[i][0]
		// srange := my_almanac.seeds[i][1]

		// for c := init; c <= init + srange; c++ {
		// 	cval := findLocation(my_almanac, c)

		// 	shortest = min(cval, shortest)
		// }
	}

	for i := 0; i < seed_len; i++ {
		seed_val := <- seed_chan[i]

		shortest = min(shortest, seed_val)
	}

	return shortest
}

func findShortestLocationByRange(my_almanac almanac, seed_range [2]int) int {
	rinit := seed_range[0]
	rrange := seed_range[1]

	shortest := findLocation(my_almanac, rinit)

	for c := rinit + 1; c < rinit + rrange; c++ {
		shortest = min(shortest, findLocation(my_almanac, c))
	}

	return shortest
}

func findLocation(my_almanac almanac, seed int) int {
	soil := getIndex(my_almanac.seed_to_soil, seed)
	fertilizer := getIndex(my_almanac.soil_to_fertilizer, soil)
	water := getIndex(my_almanac.fertilizer_to_water, fertilizer)
	light := getIndex(my_almanac.water_to_light, water)
	temperature := getIndex(my_almanac.light_to_temperature, light)
	humidity := getIndex(my_almanac.temperature_to_humidity, temperature)
	humidity_to_location := getIndex(my_almanac.humidity_to_location, humidity)

	return humidity_to_location
}

func getIndex(amap_arr []almanac_map, value int) int {
	for _, amap := range amap_arr {
		if value >= amap.source && value <= amap.source + amap.arange {
			gap := amap.dest - amap.source

			return gap + value
		}
	}

	return value
}
