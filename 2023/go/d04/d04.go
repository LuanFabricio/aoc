package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	winners []int;
	my map[int]bool;
}

func main() {
	cards := getData("assets/input1")

	fmt.Println("Cards")
	fmt.Println(cards)

	fmt.Printf("Points sum: %v\n", calcPoints(cards))
	fmt.Printf("Cards sum: %v\n", calcCards(cards))
}

func getData(filename string) []card {
	filedata, _ := os.ReadFile(filename)

	file := strings.Split(string(filedata), "\n")
	file = file[:len(file)-1]

	cards := make([]card, len(file))

	for i := 0; i < len(cards) ; i++ {
		cards_str := strings.Split(file[i], ": ")[1]

		cards_split := strings.Split(strings.TrimSpace(cards_str), " | ")
		winners := strings.Split(cards_split[0], " ")
		my := strings.Split(cards_split[1], " ")

		fmt.Println(winners)
		fmt.Println(my)

		cards[i] = card {
			winners: map_arr(winners),
		}

		cards[i].my = arr_to_map(map_arr(my))
	}

	return cards
}

func map_arr(arr []string) []int {
	num_arr := make([]int, len(arr))


	final_len := len(num_arr)
	c := 0
	for _, val := range arr {
		if len(val) == 0{
			final_len--
			continue
		}

		num_val, _ := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
		num_arr[c] = int(num_val)

		c++
	}

	return num_arr[:final_len]
}

func arr_to_map(arr []int) map[int]bool {
	result := make(map[int]bool)

	for _, num := range arr {
		result[num] = true
	}

	return result
}

func calcPoints(cards []card) int {
	total_points := 0

	for _, current_card := range cards {
		is_first_match := true

		run_points := 0

		for _, w_card := range current_card.winners {
			if current_card.my[w_card] {
				if is_first_match {
					run_points++
					is_first_match = false
				} else {
					run_points *= 2
				}
			}
		}

		total_points += run_points
		fmt.Println(run_points)
		fmt.Println(total_points)
	}

	return total_points
}

func calcCards(cards []card) int {
	copies := make(map[int]int)

	fmt.Printf("Copies: %v\n", copies)

	total_cards := len(cards)
	for card_index, current_card := range cards {
		hits := 0

		for _, w_card := range current_card.winners {
			if current_card.my[w_card] {
				hits++
			}
		}

		for i := card_index + 1; i <= card_index + hits; i++ {
			copies[i] += (copies[card_index] + 1)
		}

		fmt.Printf("Copies: %v\n", copies)
	}

	for i := 0; i < len(cards); i++ {
		total_cards += copies[i]
	}

	return total_cards
}
